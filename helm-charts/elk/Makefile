NAMESPACE := e-commerce
RELEASE := e-commerce-release
IMAGE_TAG := 8.0.0-SNAPSHOT
ELASTICSEARCH_IMAGE := docker.elastic.co/elasticsearch/elasticsearch:$(IMAGE_TAG)
DNS := elastic

init:
	helm repo add stable https://charts.helm.sh/stable
	helm repo add elastic https://helm.elastic.co
	helm repo update
	kubectl create ns $(NAMESPACE)

secrets-elastic:
	docker rm -f elastic-helm-charts-certs || true
	rm -f elastic-certificates.p12 elastic-certificate.pem elastic-certificate.crt elastic-stack-ca.p12 || true
	password=$$([ ! -z "$$ELASTIC_PASSWORD" ] && echo $$ELASTIC_PASSWORD || echo $$(docker run --rm busybox:1.31.1 /bin/sh -c "< /dev/urandom tr -cd '[:alnum:]' | head -c20")) && echo "Password for ElasticSearch: "$$password &&  \
	docker run --name elastic-helm-charts-certs -i -w /app \
		$(ELASTICSEARCH_IMAGE) \
		/bin/sh -c " \
			elasticsearch-certutil ca --out elastic-stack-ca.p12 --pass '' && \
			elasticsearch-certutil cert --name $(DNS) --dns $(DNS) --ca elastic-stack-ca.p12 --pass '' --ca-pass '' --out elastic-certificates.p12" && \
	docker cp elastic-helm-charts-certs:/usr/share/elasticsearch/elastic-certificates.p12 ./ && \
	docker rm -f elastic-helm-charts-certs && \
	openssl pkcs12 -nodes -passin pass:'' -in elastic-certificates.p12 -out elastic-certificate.pem && \
	openssl x509 -outform der -in elastic-certificate.pem -out elastic-certificate.crt && \
	kubectl -n $(NAMESPACE) create secret generic elastic-certificates --from-file=elastic-certificates.p12 && kubectl -n $(NAMESPACE) label secret elastic-certificates "app.kubernetes.io/managed-by=Helm" && kubectl -n $(NAMESPACE) annotate secret elastic-certificates "meta.helm.sh/release-name=${RELEASE}" "meta.helm.sh/release-namespace=${NAMESPACE}" && \
	kubectl -n $(NAMESPACE) create secret generic elastic-certificate-pem --from-file=elastic-certificate.pem && kubectl -n $(NAMESPACE) label secret elastic-certificate-pem "app.kubernetes.io/managed-by=Helm" && kubectl -n $(NAMESPACE) annotate secret elastic-certificate-pem "meta.helm.sh/release-name=${RELEASE}" "meta.helm.sh/release-namespace=${NAMESPACE}" && \
	kubectl -n $(NAMESPACE) create secret generic elastic-certificate-crt --from-file=elastic-certificate.crt && kubectl -n $(NAMESPACE) label secret elastic-certificate-crt "app.kubernetes.io/managed-by=Helm" && kubectl -n $(NAMESPACE) annotate secret elastic-certificate-crt "meta.helm.sh/release-name=${RELEASE}" "meta.helm.sh/release-namespace=${NAMESPACE}" && \
	kubectl -n $(NAMESPACE) create secret generic elastic-credentials  --from-literal=password=$$password --from-literal=username=elastic && kubectl -n $(NAMESPACE) label secret elastic-credentials "app.kubernetes.io/managed-by=Helm" && kubectl -n $(NAMESPACE) annotate secret elastic-credentials "meta.helm.sh/release-name=${RELEASE}" "meta.helm.sh/release-namespace=${NAMESPACE}" && \
	rm -f elastic-certificates.p12 elastic-certificate.pem elastic-certificate.crt elastic-stack-ca.p12

secrets-kibana:
	encryptionkey=$$(docker run --rm busybox:1.31.1 /bin/sh -c "< /dev/urandom tr -dc _A-Za-z0-9 | head -c50") && \
	kubectl -n $(NAMESPACE) create secret generic kibana --from-literal=encryptionkey=$$encryptionkey && kubectl -n $(NAMESPACE) label secret kibana "app.kubernetes.io/managed-by=Helm" && kubectl -n $(NAMESPACE) annotate secret kibana "meta.helm.sh/release-name=${RELEASE}" "meta.helm.sh/release-namespace=${NAMESPACE}" 

secret:
	make secrets-elastic
	make secrets-kibana

clear-certs:
	kubectl -n $(NAMESPACE) delete secrets elastic-certificate-crt elastic-certificate-pem elastic-certificates elastic-credentials kibana

install-elastic:
	helm upgrade --install elastic elastic/elasticsearch -n $(NAMESPACE) --set imageTag=$(IMAGE_TAG) -f ./values/values-elastic.yaml --version 8.5.1

install-kibana:
	helm upgrade --install kibana elastic/kibana -n $(NAMESPACE) --set imageTag=$(IMAGE_TAG) -f ./values/values-kibana.yaml --version 8.5.1

install-logstash:
	helm upgrade --install logstash elastic/logstash -n $(NAMESPACE) --set imageTag=$(IMAGE_TAG) -f ./values/values-logstash.yaml --version 8.5.1

install-filebeat:
	helm upgrade --install filebeat elastic/filebeat -n $(NAMESPACE) --set imageTag=$(IMAGE_TAG) -f ./values/values-filebeat.yaml --version 8.5.1

install:
	make install-elastic
	make install-kibana
	make install-logstash
	make install-filebeat

install-elk:
	make install-elastic
	make install-kibana
	make install-logstash

uninstall:
	helm delete -n $(NAMESPACE) elastic
	helm delete -n $(NAMESPACE) kibana
	helm delete -n $(NAMESPACE) logstash
	helm delete -n $(NAMESPACE) filebeat

deploy:
	make init
	make secret
	make install

clean:
	make clear-certs
	make uninstall