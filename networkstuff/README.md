(> means enter into a command line - probably as an admin)

This will work on all OSs that support Docker. For Windows, you will need a text editor other than notepad (like Notepad++) that isn't a Rich Text Editor (like Wordpad, Word, etc).

Honestly, for Windows... it might be easier to just download WSL2 (Linux Subsystem for Windows version 2... and run this on one of MS's Linux distros like CentOS) or install a CentOS VM. I'm not sure how to translate the bash command into a Windows command.

## Step 1. Install Docker.

For Windows/MacOS: http://www.docker.com and install "Docker Desktop" for your OS. You may need to restart after you do this.

If you use RHEL/CentOS:

\> yum --nogpgcheck -y -q -e 0 install --downloadonly --downloaddir=./rpm/docker yum-utils device-mapper-persistent-data lvm2

\> yum --nogpgcheck -y -q -e 0 install yum-utils device-mapper-persistent-data lvm2

\> yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo

\> yum --nogpgcheck -y -q -e 0 install --downloadonly --downloaddir=./rpm/docker docker-ce

\> yum --nogpgcheck -y -q -e 0 install docker-ce

\> systemctl start docker

What does docker do? It's a paravirtualization tool that lets you run a linux instance and an application inside of a 'fake OS'. The inside of the 'container' thinks your machine is a standard linux OS and the only application that is running is itself. However, you can run multiple containers together (even multiple versions of the same thing... or even multiple instances of the same version) on one machine. This basically means you can install and run Linux software on Windows or MacOS. If you're interested in this as a training topic, I can guide you through a vinette I wrote, run a live training, or let you go through one of the projects I wrote using Docker. (https://github.com/descendency/CozyNSM or https://github.com/descendency/tictactoe)

## Step 2. Download Zeek, FileBeat, Logstash, ElasticSearch, Kibana, Splunk.

\> docker pull blacktop/zeek

\> docker pull docker.elastic.co/beats/filebeat-oss:7.8.0

\> docker pull docker.elastic.co/logstash/logstash-oss:7.8.0

\> docker pull docker.elastic.co/elasticsearch/elasticsearch-oss:7.8.0-amd64

\> docker pull docker.elastic.co/kibana/kibana-oss:7.8.0

\> docker pull splunk/splunk

\> docker pull busybox

This downloads 'images' of the programs. An image is what the application needs to run (at a minimum). Think of it like a rubber stamp.

Let's make some stamps.

## Step 3. Run Zeek.

(pre-Step 3... you need to download a pcap file. I used the one we use on the board - but for obvious reasons, you will need to provide your own. malware-traffic-analysis.net has great ones)

Download the local.zeek file and put it into a directory you can easily access. (it is in this repository)

(not sure what this command is on Windows... but you need to specify the directories where local.zeek and your pcap are - ignore the right side of the ':' unless you dont want this to work.)

\> docker run --rm -v $(pwd):/pcap -v $(pwd)/local.zeek:/usr/local/zeek/share/zeek/site/local.zeek blacktop/zeek -r yourpcapsname.pcap local "Site::local_nets += { 192.168.11.0/24 }"

If done correctly, the local directory is where your zeek logs will be dumped.

## Step 4. Run ElasticSearch.

This might feel weird to start in the middle of the ELK pipeline, but you have to.

\> docker network create --driver=bridge --subnet=172.18.0.0/24 --gateway=172.18.0.1 --ipv6 --subnet=2001:3200:3200::/64 --gateway=2001:3200:3200::1 databridge

\> docker run --network databridge --ip 172.18.0.3 -itd --name es --restart=always -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch-oss:7.8.0

The 172.18.0.3 is the docker IP address. Do not change this unless you want a headache later.


## Step 5. Run Kibana.

Attach this to elasticsearch. You will need a kibana.yml that I included in the repo.

\> docker run -itd --restart=always --network databridge --ip 172.18.0.4 -v $(pwd)/kibana.yml:/usr/share/kibana/config/kibana.yml --link es:elasticsearch --name kibana -p 5601:5601 docker.elastic.co/kibana/kibana-oss:7.8.0

## Step 6. Run Logstash.

You will need my logstash.conf in the directory as well.

\> docker run --network databridge --ip 172.18.0.2 -v $(pwd)/logstash.conf:/usr/share/logstash/pipeline/logstash.conf --restart=always -itd --name logstash docker.elastic.co/logstash/logstash-oss:7.8.0

## Step 7. Run FileBeat.

This needs a filebeat.yml file.

\> docker run --network databridge -v $(pwd)/filebeat.yml:/usr/share/filebeat/filebeat.yml -itd --name filebeat -v $(pwd):/data docker.elastic.co/beats/filebeat-oss:7.8.0

\> docker exec -itu root filebeat chown filebeat:filebeat /usr/share/filebeat/filebeat.yml

\> docker restart filebeat


After a few moments, you should start to see logs populating into elasticsearch. You can see this buy running

\> curl http://127.0.0.1:9200/_cat/indices

or

Browsing to http://127.0.0.1:9200/_cat/indices

If you do not see 'logstash-...' somewhere, it isn't working. (give it a few minutes, just to be sure)

## Step 8. Delete stuff you don't need.

Now that you have loaded files into ElasticSearch and Kibana is running, you can stop/remove FileBeat and Logstash (to tidy up). You can do the same with ElasticSearch, Kibana, and Splunk when you are done.

\> docker rm -f -v filebeat logstash

## Step 9. Configure Kibana indexes.

Browse to http://127.0.0.1:5601/ > click the triple - in the top left > Stack management > Index Patterns > Create Index Pattern > index name: "logstash-\*" > next > Time Field: 'ts' > create index pattern. > triple - > 'Discover'. Bam.

## Step 10. Run Splunk.

Busybox is needed for Splunk for some reason. Not sure.

\> docker run --network databridge --restart=always -itd --name vsplunk -v /opt/splunk/etc -v /opt/splunk/var busybox

\> docker run --network databridge --ip 172.18.0.10 --restart=always -itd --name splunk --volumes-from=vsplunk -p 8000:8000 -p 9997:9997 -p 8088:8088 -p 8089:8089 -p 1514:1514 -e "SPLUNK_START_ARGS=--accept-license" -e "SPLUNK_PASSWORD=CHANGEmePLEASE" splunk/splunk

At the end of this command, you need to specify an administrator password. It must be 8+ characters, no space.

(the username is admin)

Browse to http://127.0.0.1:8000. Log in. Click 'add data' > upload > Select File (you can do them one by one or upload a zip/tar) and select your zeek logs > 'Create new index' > name it something > 'Review' > 'Submit'.


## Step 11. Completed.

If you made it this far and had no errors, tell your friends/peers how easy this was, how little you had to do, and how easy it is for them to spin up their own training.

If you didn't make it this far and you just skipped ahead hoping for magic - please ask me questions. I can help with Docker, Zeek, Filebeat, Logstash, ElasticSearch, Kibana, Splunk, and whatever OS you are using... even Windows.

## Step ???. Clean up.

After you are done and no longer want splunk, kibana, AND es...

\> docker rm -f -v $(docker -q)

or

\> docker rm -f -v es kibana splunk

If you would like to free up some disk space from the images you downloaded...

\> docker rmi blacktop/zeek busybox splunk/splunk docker.elastic.co/logstash/logstash-oss:7.8.0 docker.elastic.co/kibana/kibana-oss:7.8.0 docker.elastic.co/elasticsearch/elasticsearch-oss:7.8.0 docker.elastic.co/beats/filebeat-oss:7.8.0

or

\> docker rmi $(docker images -q)
