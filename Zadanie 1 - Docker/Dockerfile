FROM ubuntu

ENV TZ=Europe/Warsaw

RUN apt update -y && \
	apt install tzdata -y && \
	ln -fs /usr/share/zoneinfo/Europe/Warsaw /etc/localtime && \
	dpkg-reconfigure -f noninteractive tzdata && \
	apt-get install software-properties-common -y && \
	add-apt-repository ppa:deadsnakes/ppa -y && \
	apt update -y && \
	apt install python3.8 -y
	
CMD python3.8 -V
