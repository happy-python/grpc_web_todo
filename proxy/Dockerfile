FROM envoyproxy/envoy:v1.8.0
COPY envoy.yaml /etc/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy.yaml
EXPOSE 9090