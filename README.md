# Statsd Notifier
> This app contains just one POST request. When this request is received, the metric is sent to the Statsd over UDP. 

To start the application, you need to create a configuration file `config.yaml`.
```yaml
server:
  debug_mode: True
  host: 127.0.0.1
  port: 8077
statsd:
  mocked: False
  host: 127.0.0.1
  port: 8125
```
Where it is necessary to specify server and statsd endpoint options. 
Name of `Counters` metric is fixed `production.fqdn.statsd.post.request.counter"`.

Then use the command to launch the application:
```bash
    make
```
To run the application without tests, you can use the command:
```bash
    make run-app
```

## Testing
After starting the application to simulate traffic, you can use the command:
```bash
    make send-reqests
```
As a result, we can see something similar
![Image of Yaktocat](https://github.com/serboox/statsd-notifier/blob/master/screen_for_example.png?raw=true)
