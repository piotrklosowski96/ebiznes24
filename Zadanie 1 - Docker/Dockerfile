FROM zadanie-1:wymaganie-2

WORKDIR /tmp

RUN wget https://services.gradle.org/distributions/gradle-8.7-bin.zip && \
	mkdir /opt/gradle && \
	unzip -d /usr/local/bin/gradle gradle-8.7-bin.zip && \
	rm -rf /tmp/*
	
ENV PATH=$PATH:/usr/local/bin/gradle/gradle-8.7/bin
ENV JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64
ENV PATH=$PATH:$JAVA_HOME/bin

WORKDIR /home

COPY ./gradle-project ./gradle-project

RUN	cd ./gradle-project && \
	gradle check && \
	gradle build

WORKDIR /home/gradle-project

CMD gradle run
