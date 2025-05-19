# Auth Service Core API

## Usage
The micro-service comes with both a public (port 8080) and private (port 8081) api server.
This service is supposed to be set up behind a load balancer, allowing public access only to the public server while allowing other internal services to communicate with the private server.

**Public Endpoint Examples**
<ul style="padding:0; margin:0; margin-top:-10px;">
<li>
    POST auth-service:8080/login HTTP/1.1
    <p style="margin:5px; padding:0; padding-left:20px;">
        Public endpoint to get auth token with username-password
    </p>
</li>
<li>
    POST auth-service:8080/signup HTTP/1.1
    <p style="margin:5px; padding:0; padding-left:20px;">
        Public endpoint to register a new user
    </p>
</li>
</ul>
<br/>

**Private Endpoint Examples**
<ul style="padding:0; margin:0;margin-top:-10px;">
<li>
    GET auth-service:8081/session/$id/info HTTP/1.1
    <p style="margin:5px; padding:0; padding-left:20px;">
        Get private info about the session and the associated user
    </p>
</li>
<li>
    GET auth-service:8081/api_key/$key/info HTTP/1.1
    <p style="margin:5px; padding:0; padding-left:20px;">
        Get private info about an api key and the associated user
    </p>
</li>
</ul>
<br/>


Full Swagger docs can be found [here](#)