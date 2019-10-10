# Configuring an HTTP Proxy for obsutil<a name="EN-US_TOPIC_0183789343"></a>

You can configure an HTTP proxy in either of the following ways:

Method 1: Set the  **proxyUrl**  parameter in the  **.obsutilconfig**  file. Example:  **proxyUrl=http://username:password@your-proxy:8080**;

Method 2: Use the system environment variable  **HTTPS\_PROXY**  or  **HTTP\_PROXY**. Example:  **HTTPS\_PROXY=http://username:password@your-proxy:8080**.

>![](public_sys-resources/icon-note.gif) **NOTE:**   
>-   HTTP proxy format:  **http://\[_Username_:_Password_@\]_Proxy server address_:_Port number_**. The  _Username_  and  _Password_  are optional.  
>-   The  **proxyUrl**  parameter and system environment variables are in the following priority order:  **proxyUrl**  \>  **HTTPS\_PROXY**  \>  **HTTP\_PROXY**.  
>-   The user name and password cannot contain colons \(:\) and at signs \(@\), which will result in parsing errors.  

