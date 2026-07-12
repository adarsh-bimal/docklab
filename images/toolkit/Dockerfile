FROM kalilinux/kali-rolling

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y \
        nmap \
        netcat-openbsd \
        curl \
        iputils-ping \
        wget \
        httpie \
        jq \
        ffuf \
        feroxbuster \
        gobuster \
        whatweb \
        nikto \
        sqlmap \
        hydra \
        python3 \
        python3-pip \
        git \
        vim \
        tmux \
        tree \
        dnsutils \
        iproute2 \
        net-tools \
        metasploit-framework \
        exploitdb \
        seclists \
        wordlists \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*



RUN git clone https://github.com/ticarpi/jwt_tool.git /opt/jwt-tool && \
    pip3 install --break-system-packages -r /opt/jwt-tool/requirements.txt && \
    ln -s /opt/jwt-tool/jwt_tool.py /usr/local/bin/jwt-tool

RUN ln -sf /usr/share/seclists /opt/SecLists

RUN if [ -f /usr/share/wordlists/rockyou.txt.gz ]; then \
        gunzip /usr/share/wordlists/rockyou.txt.gz; \
    fi

CMD ["sleep", "infinity"]