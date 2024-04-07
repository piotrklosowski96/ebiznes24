FROM sbtscala/scala-sbt:eclipse-temurin-jammy-21.0.2_13_1.9.9_2.13.13 as builder

WORKDIR /home/sbtuser
COPY ./scala-project scala-project

WORKDIR /home/sbtuser/scala-project
RUN sbt assembly


FROM eclipse-temurin:21.0.2_13-jre-jammy

ARG NGROK_TOKEN

RUN curl -s https://ngrok-agent.s3.amazonaws.com/ngrok.asc | tee /etc/apt/trusted.gpg.d/ngrok.asc >/dev/null && \
	echo "deb https://ngrok-agent.s3.amazonaws.com buster main" | tee /etc/apt/sources.list.d/ngrok.list && \
	apt update && \
	apt install ngrok && \
	ngrok config add-authtoken ${NGROK_TOKEN}
	
RUN	mkdir /opt/zadanie2-assembly

COPY --from=builder /home/sbtuser/scala-project/target/scala-2.13/zadanie2-assembly-1.0.jar /opt/zadanie2-assembly
COPY ./scala-project/start.sh /opt/zadanie2-assembly/start.sh

ENTRYPOINT /bin/bash -c /opt/zadanie2-assembly/start.sh