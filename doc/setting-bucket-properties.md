# Setting Bucket Properties<a name="EN-US_TOPIC_0149246212"></a>

## Function<a name="section1479112110815"></a>

You can use this command to set the properties of a bucket, such as storage classes and  access policies.

## Command Line Structure<a name="section1220752192216"></a>

-   In Windows

    ```
    obsutil chattri obs://bucket [-sc=xxx] [-acl=xxx] [-aclXml=xxx] [-config=xxx]
    ```

-   In Linux or macOS

    ```
    ./obsutil chattri obs://bucket [-sc=xxx] [-acl=xxx] [-aclXml=xxx] [-config=xxx]
    ```


## Parameter Description<a name="section6559191102418"></a>

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="17.82178217821782%" id="mcps1.1.4.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="25.742574257425744%" id="mcps1.1.4.1.2"><p id="p15137125919108"><a name="p15137125919108"></a><a name="p15137125919108"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="56.43564356435643%" id="mcps1.1.4.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row108328217449"><td class="cellrowborder" valign="top" width="17.82178217821782%" headers="mcps1.1.4.1.1 "><p id="p64495172515"><a name="p64495172515"></a><a name="p64495172515"></a>bucket</p>
</td>
<td class="cellrowborder" valign="top" width="25.742574257425744%" headers="mcps1.1.4.1.2 "><p id="p154316502519"><a name="p154316502519"></a><a name="p154316502519"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="56.43564356435643%" headers="mcps1.1.4.1.3 "><p id="p17425512259"><a name="p17425512259"></a><a name="p17425512259"></a>Bucket name</p>
</td>
</tr>
<tr id="row8533319194211"><td class="cellrowborder" valign="top" width="17.82178217821782%" headers="mcps1.1.4.1.1 "><p id="p19533119154211"><a name="p19533119154211"></a><a name="p19533119154211"></a>sc</p>
</td>
<td class="cellrowborder" valign="top" width="25.742574257425744%" headers="mcps1.1.4.1.2 "><p id="p4533191944218"><a name="p4533191944218"></a><a name="p4533191944218"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.43564356435643%" headers="mcps1.1.4.1.3 "><p id="p86547153813"><a name="p86547153813"></a><a name="p86547153813"></a>Default storage class of the bucket. Possible values are:</p>
<a name="ul175651814214"></a><a name="ul175651814214"></a><ul id="ul175651814214"><li><strong id="b717243073212"><a name="b717243073212"></a><a name="b717243073212"></a>standard</strong>: OBS Standard, which features low access latency and high throughput, and is applicable to storing frequently accessed data (multiple accesses per month averagely) or data that is smaller than 1 MB</li><li><strong id="b161621451521"><a name="b161621451521"></a><a name="b161621451521"></a>warm</strong>: OBS Warm. It is applicable to storing semi-frequently accessed (less than 12 times a year averagely) data that requires quick response.</li><li><strong id="b11585340210"><a name="b11585340210"></a><a name="b11585340210"></a>cold</strong>: OBS Cold. It is secure, durable, and inexpensive, and applicable to archiving rarely-accessed (once a year averagely) data.</li></ul>
</td>
</tr>
<tr id="row184781221162513"><td class="cellrowborder" valign="top" width="17.82178217821782%" headers="mcps1.1.4.1.1 "><p id="p12478192122514"><a name="p12478192122514"></a><a name="p12478192122514"></a>acl</p>
</td>
<td class="cellrowborder" valign="top" width="25.742574257425744%" headers="mcps1.1.4.1.2 "><p id="p82428286353"><a name="p82428286353"></a><a name="p82428286353"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.43564356435643%" headers="mcps1.1.4.1.3 "><p id="p16811512123619"><a name="p16811512123619"></a><a name="p16811512123619"></a>Access control policies that can be specified for buckets. Possible values are:</p>
<a name="ul1273864144412"></a><a name="ul1273864144412"></a><ul id="ul1273864144412"><li>private</li><li>public-read</li><li>public-read-write</li></ul>
<div class="note" id="note1790113183525"><a name="note1790113183525"></a><a name="note1790113183525"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p92982882916"><a name="p92982882916"></a><a name="p92982882916"></a>The preceding three values indicate private read and write, public read, and public read and write.</p>
</div></div>
</td>
</tr>
<tr id="row25713416178"><td class="cellrowborder" valign="top" width="17.82178217821782%" headers="mcps1.1.4.1.1 "><p id="p1258123413176"><a name="p1258123413176"></a><a name="p1258123413176"></a>aclXml</p>
</td>
<td class="cellrowborder" valign="top" width="25.742574257425744%" headers="mcps1.1.4.1.2 "><p id="p75818346179"><a name="p75818346179"></a><a name="p75818346179"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.43564356435643%" headers="mcps1.1.4.1.3 "><p id="p2058173418170"><a name="p2058173418170"></a><a name="p2058173418170"></a>Access control policy of the bucket, in XML format.</p>
<pre class="screen" id="screen78378416219"><a name="screen78378416219"></a><a name="screen78378416219"></a>&lt;AccessControlPolicy&gt;
    &lt;Owner&gt;
        &lt;ID&gt;<em id="i131811237783"><a name="i131811237783"></a><a name="i131811237783"></a>ownerid</em>&lt;/ID&gt;
    &lt;/Owner&gt;
    &lt;AccessControlList&gt;
        &lt;Grant&gt;
            &lt;Grantee&gt;
                &lt;ID&gt;<em id="i20604651896"><a name="i20604651896"></a><a name="i20604651896"></a>userid</em>&lt;/ID&gt;
            &lt;/Grantee&gt;
            &lt;Permission&gt;<em id="i786413238102"><a name="i786413238102"></a><a name="i786413238102"></a>[WRITE|WRITE_ACP|<em id="i198641723191017"><a name="i198641723191017"></a><a name="i198641723191017"></a>READ</em>|READ_ACP|FULL_CONTROL]</em>&lt;/Permission&gt;
        &lt;/Grant&gt;
        &lt;Grant&gt;
            &lt;Grantee&gt;
                &lt;Canned&gt;Everyone&lt;/Canned&gt;
            &lt;/Grantee&gt;
            &lt;Permission&gt;<em id="i287517813170"><a name="i287517813170"></a><a name="i287517813170"></a>[WRITE|WRITE_ACP|<em id="i18751384176"><a name="i18751384176"></a><a name="i18751384176"></a>READ</em>|READ_ACP|FULL_CONTROL]</em>&lt;/Permission&gt;
        &lt;/Grant&gt;
    &lt;/AccessControlList&gt;
&lt;/AccessControlPolicy&gt;</pre>
<div class="note" id="note8740143916331"><a name="note8740143916331"></a><a name="note8740143916331"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul11591446103319"></a><a name="ul11591446103319"></a><ul id="ul11591446103319"><li><strong id="b12293102015418"><a name="b12293102015418"></a><a name="b12293102015418"></a>Owner</strong>: Optional. Specify the bucket owner's ID.</li><li>In <strong id="b139036161165"><a name="b139036161165"></a><a name="b139036161165"></a>AccessControlList</strong>, the <strong id="b1856172415618"><a name="b1856172415618"></a><a name="b1856172415618"></a>Grant</strong> field contains the authorized users. <strong id="b1515515681315"><a name="b1515515681315"></a><a name="b1515515681315"></a>Grantee</strong> specifies the IDs of authorized users. <strong id="b129881752101320"><a name="b129881752101320"></a><a name="b129881752101320"></a>Canned</strong> specifies the authorized user group (currently, only <strong id="b8965141115154"><a name="b8965141115154"></a><a name="b8965141115154"></a>Everyone</strong> is supported).</li><li>The following permissions can be granted: WRITE (write), WRITE_ACP (write ACL), READ (read), READ_ACP (read ACL), and FULL_CONTROL (full control).</li></ul>
</div></div>
<div class="notice" id="note14101733103315"><a name="note14101733103315"></a><a name="note14101733103315"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><p id="p11101733143314"><a name="p11101733143314"></a><a name="p11101733143314"></a>Because angle brackets (&lt;) and (&gt;) are unavoidably included in the parameter value, you must use quotation marks to enclose them for escaping when running the command. Use single quotation marks for Linux or macOS and quotation marks for Windows.</p>
</div></div>
</td>
</tr>
<tr id="row10314444123413"><td class="cellrowborder" valign="top" width="17.82178217821782%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="25.742574257425744%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.43564356435643%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

>![](public_sys-resources/icon-note.gif) **NOTE:**   
>Only  **sc**,  **acl**, or  **aclXml**  can be set for each command.  

## Running Example<a name="section15899161919244"></a>

-   Take the Windows OS as an example. Run the  **obsutil chattri obs://bucket-test -acl=private**  command to change the access control policy of the bucket to private read and write.

```
obsutil chattri obs://bucket-test -acl=private

Set the acl of bucket [bucket-test] to [private] successfully, request id [04050000016836C5DA6FB21F14A2A0C0]
```

