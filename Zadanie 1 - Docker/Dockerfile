FROM ubuntu:22.04

WORKDIR /tmp

RUN apt update && \
	apt install openjdk-8-jdk wget unzip -y && \
	wget https://github.com/JetBrains/kotlin/releases/download/v1.9.23/kotlin-compiler-1.9.23.zip && \
	unzip ./kotlin-compiler-1.9.23.zip -d /usr/local/bin && \
	rm -rf /tmp/*
	
ENV PATH=$PATH:/usr/local/bin/kotlinc/bin

CMD java -version && kotlinc -version
