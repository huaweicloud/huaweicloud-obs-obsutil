# Creating a Bucket<a name="EN-US_TOPIC_0141181362"></a>

## Function<a name="section1479112110815"></a>

You can use this command to create a bucket. A bucket name must be unique in OBS. One account can create a maximum of 100 buckets.

>![](public_sys-resources/icon-note.gif) **NOTE:**   
>If you create a bucket and name it the same as an existing one in the same account and region, no error will be reported and status code 200 is returned. The bucket properties comply with those set in the first creation request. In other cases, creating a namesake bucket will receive the status code 409, indicating that the bucket already exists.  

## Command Line Structure<a name="section1220752192216"></a>

-   In Windows

    ```
    obsutil mb obs://bucket [-acl=xxx] [-sc=xxx] [-location=xxx] [-config=xxx]
    ```

-   In Linux or macOS

    ```
    ./obsutil mb obs://bucket [-acl=xxx] [-sc=xxx] [-location=xxx] [-config=xxx] 
    ```


## Parameter Description<a name="section6559191102418"></a>

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="13%" id="mcps1.1.4.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="22%" id="mcps1.1.4.1.2"><p id="p15137125919108"><a name="p15137125919108"></a><a name="p15137125919108"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="65%" id="mcps1.1.4.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row108328217449"><td class="cellrowborder" valign="top" width="13%" headers="mcps1.1.4.1.1 "><p id="p64495172515"><a name="p64495172515"></a><a name="p64495172515"></a>bucket</p>
</td>
<td class="cellrowborder" valign="top" width="22%" headers="mcps1.1.4.1.2 "><p id="p154316502519"><a name="p154316502519"></a><a name="p154316502519"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="65%" headers="mcps1.1.4.1.3 "><p id="p17425512259"><a name="p17425512259"></a><a name="p17425512259"></a>Bucket name</p>
<div class="note" id="note578772312297"><a name="note578772312297"></a><a name="note578772312297"></a><span class="notetitle"> NOTE: </span><div class="notebody"><div class="p" id="p1088423931013"><a name="p1088423931013"></a><a name="p1088423931013"></a>A bucket name must comply with the following rules:<a name="ul18722561087"></a><a name="ul18722561087"></a><ul id="ul18722561087"><li>Contains 3 to 63 characters, including lowercase letters, digits, hyphens (-), and periods (.), and starts with a digit or letter.</li><li>Cannot be an IP address.</li><li>Cannot start or end with a hyphen (-) or period (.).</li><li>Cannot contain two consecutive periods (.), for example, <strong id="b842352706191036"><a name="b842352706191036"></a><a name="b842352706191036"></a>my..bucket</strong>.</li><li>Cannot contain periods (.) and hyphens (-) adjacent to each other, for example, <strong id="b842352706191112"><a name="b842352706191112"></a><a name="b842352706191112"></a>my-.bucket</strong> or <strong id="b842352706191116"><a name="b842352706191116"></a><a name="b842352706191116"></a>my.-bucket</strong>.</li></ul>
</div>
</div></div>
</td>
</tr>
<tr id="row19509222417"><td class="cellrowborder" valign="top" width="13%" headers="mcps1.1.4.1.1 "><p id="p159501222414"><a name="p159501222414"></a><a name="p159501222414"></a>acl</p>
</td>
<td class="cellrowborder" valign="top" width="22%" headers="mcps1.1.4.1.2 "><p id="p4950172216419"><a name="p4950172216419"></a><a name="p4950172216419"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65%" headers="mcps1.1.4.1.3 "><p id="p16811512123619"><a name="p16811512123619"></a><a name="p16811512123619"></a>Access control policies that can be specified when creating a bucket. Possible values are:</p>
<a name="ul1273864144412"></a><a name="ul1273864144412"></a><ul id="ul1273864144412"><li>private</li><li>public-read</li><li>public-read-write</li></ul>
<div class="note" id="note112972817292"><a name="note112972817292"></a><a name="note112972817292"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p92982882916"><a name="p92982882916"></a><a name="p92982882916"></a>The preceding three values indicate private read and write, public read, and public read and write.</p>
</div></div>
</td>
</tr>
<tr id="row8533319194211"><td class="cellrowborder" valign="top" width="13%" headers="mcps1.1.4.1.1 "><p id="p19533119154211"><a name="p19533119154211"></a><a name="p19533119154211"></a>sc</p>
</td>
<td class="cellrowborder" valign="top" width="22%" headers="mcps1.1.4.1.2 "><p id="p4533191944218"><a name="p4533191944218"></a><a name="p4533191944218"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65%" headers="mcps1.1.4.1.3 "><p id="p86547153813"><a name="p86547153813"></a><a name="p86547153813"></a>Default bucket storage class that can be specified when creating a bucket. Possible values are:</p>
<a name="ul175651814214"></a><a name="ul175651814214"></a><ul id="ul175651814214"><li><strong id="b10635191912181"><a name="b10635191912181"></a><a name="b10635191912181"></a>standard</strong>: OBS Standard, which features low access latency and high throughput, and is applicable to storing frequently accessed data (multiple accesses per month averagely) or data that is smaller than 1 MB</li><li><strong id="b132258202416"><a name="b132258202416"></a><a name="b132258202416"></a>warm</strong>: OBS Warm. It is applicable to storing semi-frequently accessed (less than 12 times a year averagely) data that requires quick response.</li><li><strong id="b124791856141313"><a name="b124791856141313"></a><a name="b124791856141313"></a>cold</strong>: OBS Cold. It is secure, durable, and inexpensive, and applicable to archiving rarely-accessed (once a year averagely) data.</li></ul>
</td>
</tr>
<tr id="row116471725174219"><td class="cellrowborder" valign="top" width="13%" headers="mcps1.1.4.1.1 "><p id="p16475251428"><a name="p16475251428"></a><a name="p16475251428"></a>location</p>
</td>
<td class="cellrowborder" valign="top" width="22%" headers="mcps1.1.4.1.2 "><p id="p62572833155512"><a name="p62572833155512"></a><a name="p62572833155512"></a>Mandatory unless the region where the OBS service resides is not the default region (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65%" headers="mcps1.1.4.1.3 "><p id="p5359820175320"><a name="p5359820175320"></a><a name="p5359820175320"></a>Region where the bucket resides.</p>
</td>
</tr>
<tr id="row03949318313"><td class="cellrowborder" valign="top" width="13%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="22%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="65%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Running Example<a name="section15899161919244"></a>

-   Take the Windows OS as an example. Run the  **obsutil mb obs://bucket-test**  command to create a bucket. The creation is successful.

```
obsutil mb obs://bucket-test

Create bucket [bucket-test] successfully, request id [0000016979E1D2EA860BB5E80A6B8FCC]
```

-   Take the Windows OS as an example. Run the  **obsutil mb obs://bucket001**  command to create a namesake bucket. The creation fails.

```
obsutil mb obs://bucket001

Create bucket [bucket001] failed, http status [409], error code [BucketAlreadyExists], error message [The requested bucket name is not available. The bucket namespace is shared by all users of the system. Please select a different name andtry again.], request id [04030000016757F31A0333281A6B1E92]
```

