# Setting Object Properties<a name="EN-US_TOPIC_0149246213"></a>

## Function<a name="section1479112110815"></a>

You can use this command to set properties of an object or set properties of objects in batches by a specified object name prefix.

>![](public_sys-resources/icon-note.gif) **NOTE:**   
>You can set storage classes only for buckets whose version is 3.0.  

## Command Line Structure<a name="section1220752192216"></a>

-   In Windows
    -   Setting properties of a single object

        ```
        obsutil chattri obs://bucket/key [-sc=xxx] [-acl=xxx] [-aclXml=xxx] [-versionId=xxx] [-fr] [-o=xxx] [-config=xxx]
        ```

    -   Setting properties of objects in batches

        ```
        obsutil chattri obs://bucket[/key] -r [-f] [-v] [-sc=xxx] [-acl=xxx] [-aclXml=xxx] [-o=xxx] [-j=1] [-config=xxx]
        ```


-   In Linux or macOS
    -   Setting properties of a single object

        ```
        ./obsutil chattri obs://bucket/key [-sc=xxx] [-acl=xxx] [-aclXml=xxx] [-versionId=xxx] [-fr] [-o=xxx] [-config=xxx]
        ```

    -   Setting properties of objects in batches

        ```
        ./obsutil chattri obs://bucket[/key] -r [-f] [-v] [-sc=xxx] [-acl=xxx] [-aclXml=xxx] [-o=xxx] [-j=1] [-config=xxx]
        ```



## Parameter Description<a name="section6559191102418"></a>

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="16%" id="mcps1.1.4.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="24%" id="mcps1.1.4.1.2"><p id="p15137125919108"><a name="p15137125919108"></a><a name="p15137125919108"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="60%" id="mcps1.1.4.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row108328217449"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p64495172515"><a name="p64495172515"></a><a name="p64495172515"></a>bucket</p>
</td>
<td class="cellrowborder" valign="top" width="24%" headers="mcps1.1.4.1.2 "><p id="p154316502519"><a name="p154316502519"></a><a name="p154316502519"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="60%" headers="mcps1.1.4.1.3 "><p id="p17425512259"><a name="p17425512259"></a><a name="p17425512259"></a>Bucket name</p>
</td>
</tr>
<tr id="row16373619121810"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p737411911186"><a name="p737411911186"></a><a name="p737411911186"></a>key</p>
</td>
<td class="cellrowborder" valign="top" width="24%" headers="mcps1.1.4.1.2 "><p id="p8166018195614"><a name="p8166018195614"></a><a name="p8166018195614"></a>Mandatory for setting properties of an object.</p>
<p id="p741812412475"><a name="p741812412475"></a><a name="p741812412475"></a>Optional for setting properties of objects in batches.</p>
</td>
<td class="cellrowborder" valign="top" width="60%" headers="mcps1.1.4.1.3 "><p id="p1864718233461"><a name="p1864718233461"></a><a name="p1864718233461"></a>Indicates the name of the object whose properties are to be set, or the name prefix of objects whose properties are to be set in batches.</p>
<div class="note" id="note12611558164215"><a name="note12611558164215"></a><a name="note12611558164215"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p1061114584426"><a name="p1061114584426"></a><a name="p1061114584426"></a>If this parameter is left blank during batch operation, properties of all objects in the bucket are set.</p>
</div></div>
</td>
</tr>
<tr id="row8533319194211"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p19533119154211"><a name="p19533119154211"></a><a name="p19533119154211"></a>sc</p>
</td>
<td class="cellrowborder" valign="top" width="24%" headers="mcps1.1.4.1.2 "><p id="p4533191944218"><a name="p4533191944218"></a><a name="p4533191944218"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="60%" headers="mcps1.1.4.1.3 "><p id="p86547153813"><a name="p86547153813"></a><a name="p86547153813"></a>Storage classes of objects. Possible values are:</p>
<a name="ul175651814214"></a><a name="ul175651814214"></a><ul id="ul175651814214"><li><strong id="b1556994153417"><a name="b1556994153417"></a><a name="b1556994153417"></a>standard</strong>: OBS Standard, which features low access latency and high throughput, and is applicable to storing frequently accessed data (multiple accesses per month averagely) or data that is smaller than 1 MB</li><li><strong id="b11414933033"><a name="b11414933033"></a><a name="b11414933033"></a>warm</strong>: OBS Warm. It is applicable to storing semi-frequently accessed (less than 12 times a year averagely) data that requires quick response.</li><li><strong id="b661273810313"><a name="b661273810313"></a><a name="b661273810313"></a>cold</strong>: OBS Cold. It is secure, durable, and inexpensive, and applicable to archiving rarely-accessed (once a year averagely) data.</li></ul>
<div class="note" id="note16713103015307"><a name="note16713103015307"></a><a name="note16713103015307"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p28791713112514"><a name="p28791713112514"></a><a name="p28791713112514"></a>For an object whose storage class is <strong id="b189301154153519"><a name="b189301154153519"></a><a name="b189301154153519"></a>cold</strong>, restore the object first and then set its storage class. To restore an object, see <a href="restoring-objects-from-obs-cold.md">Restoring Objects from OBS Cold</a>.</p>
</div></div>
</td>
</tr>
<tr id="row490772910240"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p990817298248"><a name="p990817298248"></a><a name="p990817298248"></a>acl</p>
</td>
<td class="cellrowborder" valign="top" width="24%" headers="mcps1.1.4.1.2 "><p id="p82428286353"><a name="p82428286353"></a><a name="p82428286353"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="60%" headers="mcps1.1.4.1.3 "><p id="p16811512123619"><a name="p16811512123619"></a><a name="p16811512123619"></a>Access control policies that can be specified for objects. Possible values are:</p>
<a name="ul1273864144412"></a><a name="ul1273864144412"></a><ul id="ul1273864144412"><li>private</li><li>public-read</li><li>public-read-write</li><li>bucket-owner-full-control</li></ul>
<div class="note" id="note1790113183525"><a name="note1790113183525"></a><a name="note1790113183525"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p92982882916"><a name="p92982882916"></a><a name="p92982882916"></a>The preceding four values indicate private read and write, public read, public read and write, and bucket owner full control.</p>
</div></div>
</td>
</tr>
<tr id="row32753484186"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p1258123413176"><a name="p1258123413176"></a><a name="p1258123413176"></a>aclXml</p>
</td>
<td class="cellrowborder" valign="top" width="24%" headers="mcps1.1.4.1.2 "><p id="p75818346179"><a name="p75818346179"></a><a name="p75818346179"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="60%" headers="mcps1.1.4.1.3 "><p id="p2058173418170"><a name="p2058173418170"></a><a name="p2058173418170"></a>Access control policy of the bucket, in XML format.</p>
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
<div class="note" id="note8740143916331"><a name="note8740143916331"></a><a name="note8740143916331"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul11591446103319"></a><a name="ul11591446103319"></a><ul id="ul11591446103319"><li><strong id="b538215791411"><a name="b538215791411"></a><a name="b538215791411"></a>Owner</strong>: Optional. Specify the object owner's ID.</li><li>In <strong id="b10739131191420"><a name="b10739131191420"></a><a name="b10739131191420"></a>AccessControlList</strong>, the <strong id="b1774018116148"><a name="b1774018116148"></a><a name="b1774018116148"></a>Grant</strong> field contains the authorized users. <strong id="b1741201121410"><a name="b1741201121410"></a><a name="b1741201121410"></a>Grantee</strong> specifies the IDs of authorized users. <strong id="b674151161414"><a name="b674151161414"></a><a name="b674151161414"></a>Canned</strong> specifies the authorized user group (currently, only <strong id="b0742211121415"><a name="b0742211121415"></a><a name="b0742211121415"></a>Everyone</strong> is supported).</li><li>The following permissions can be granted: WRITE (write), WRITE_ACP (write ACL), READ (read), READ_ACP (read ACL), and FULL_CONTROL (full control).</li></ul>
</div></div>
<div class="notice" id="note14101733103315"><a name="note14101733103315"></a><a name="note14101733103315"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><p id="p11101733143314"><a name="p11101733143314"></a><a name="p11101733143314"></a>Because angle brackets (&lt;) and (&gt;) are unavoidably included in the parameter value, you must use quotation marks to enclose them for escaping when running the command. Use single quotation marks for Linux or macOS and quotation marks for Windows.</p>
</div></div>
</td>
</tr>
<tr id="row1141710101913"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p174267194117"><a name="p174267194117"></a><a name="p174267194117"></a>versionId</p>
</td>
<td class="cellrowborder" valign="top" width="24%" headers="mcps1.1.4.1.2 "><p id="p774216710415"><a name="p774216710415"></a><a name="p774216710415"></a>Optional for setting properties of an object (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="60%" headers="mcps1.1.4.1.3 "><p id="p874257164113"><a name="p874257164113"></a><a name="p874257164113"></a>Version ID of the object whose properties are to be set</p>
</td>
</tr>
<tr id="row12415113193"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p8517143114816"><a name="p8517143114816"></a><a name="p8517143114816"></a>fr</p>
</td>
<td class="cellrowborder" valign="top" width="24%" headers="mcps1.1.4.1.2 "><p id="p1965455113483"><a name="p1965455113483"></a><a name="p1965455113483"></a>Optional for setting properties of an object (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="60%" headers="mcps1.1.4.1.3 "><p id="p1951718433487"><a name="p1951718433487"></a><a name="p1951718433487"></a>Generates an operation result list when setting properties of an object.</p>
</td>
</tr>
<tr id="row18415121181916"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p9369940145311"><a name="p9369940145311"></a><a name="p9369940145311"></a>f</p>
</td>
<td class="cellrowborder" valign="top" width="24%" headers="mcps1.1.4.1.2 "><p id="p19278143834816"><a name="p19278143834816"></a><a name="p19278143834816"></a>Optional when setting properties of objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="60%" headers="mcps1.1.4.1.3 "><p id="p972417362487"><a name="p972417362487"></a><a name="p972417362487"></a>Runs in force mode.</p>
</td>
</tr>
<tr id="row1441461141918"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p71594425537"><a name="p71594425537"></a><a name="p71594425537"></a>r</p>
</td>
<td class="cellrowborder" valign="top" width="24%" headers="mcps1.1.4.1.2 "><p id="p1273555018488"><a name="p1273555018488"></a><a name="p1273555018488"></a>Mandatory when setting properties of objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="60%" headers="mcps1.1.4.1.3 "><p id="p1173519509481"><a name="p1173519509481"></a><a name="p1173519509481"></a>Sets properties of objects in batches based on a specified object name prefix.</p>
</td>
</tr>
<tr id="row12414181121915"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p373223134117"><a name="p373223134117"></a><a name="p373223134117"></a>v</p>
</td>
<td class="cellrowborder" valign="top" width="24%" headers="mcps1.1.4.1.2 "><p id="p101701338485"><a name="p101701338485"></a><a name="p101701338485"></a>Optional when setting properties of objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="60%" headers="mcps1.1.4.1.3 "><p id="p1073202316412"><a name="p1073202316412"></a><a name="p1073202316412"></a>Sets properties of versions of objects in batches based on a specified object name prefix.</p>
</td>
</tr>
<tr id="row1241412161915"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p138574472533"><a name="p138574472533"></a><a name="p138574472533"></a>o</p>
</td>
<td class="cellrowborder" valign="top" width="24%" headers="mcps1.1.4.1.2 "><p id="p817213354812"><a name="p817213354812"></a><a name="p817213354812"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="60%" headers="mcps1.1.4.1.3 "><p id="p4924117181112"><a name="p4924117181112"></a><a name="p4924117181112"></a>Indicates the folder where operation result lists reside. After the command is executed, result lists (including success and failure files) are generated in the folder. The default value is <strong id="b132234473221"><a name="b132234473221"></a><a name="b132234473221"></a>.obsutil_output</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note349481802414"><a name="note349481802414"></a><a name="note349481802414"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul101190347408"></a><a name="ul101190347408"></a><ul id="ul101190347408"><li>The naming rule for result lists is as follows: <strong id="b9923175232616"><a name="b9923175232616"></a><a name="b9923175232616"></a>chattri_{succeed  | failed}_report_</strong><em id="i17925205219267"><a name="i17925205219267"></a><a name="i17925205219267"></a>time</em><strong id="b392511527263"><a name="b392511527263"></a><a name="b392511527263"></a>_TaskId.txt</strong><p id="li412103664016p0"><a name="li412103664016p0"></a><a name="li412103664016p0"></a>By default, the maximum size of a single result list is 30 MB and the maximum number of result lists that can be retained is 1024. You can set the maximum size and number by configuring <strong id="b790816547339"><a name="b790816547339"></a><a name="b790816547339"></a>recordMaxLogSize</strong> and <strong id="b29081154153319"><a name="b29081154153319"></a><a name="b29081154153319"></a>recordBackups</strong> in the configuration file.</p>
</li></ul>
</div></div>
</td>
</tr>
<tr id="row541431131917"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p87050508534"><a name="p87050508534"></a><a name="p87050508534"></a>j</p>
</td>
<td class="cellrowborder" valign="top" width="24%" headers="mcps1.1.4.1.2 "><p id="p19177173320483"><a name="p19177173320483"></a><a name="p19177173320483"></a>Optional when setting properties of objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="60%" headers="mcps1.1.4.1.3 "><p id="p2549177181312"><a name="p2549177181312"></a><a name="p2549177181312"></a>Indicates the maximum number of concurrent tasks for setting object properties in batches. The default value is the value of <strong id="b1268754012466"><a name="b1268754012466"></a><a name="b1268754012466"></a>defaultJobs</strong> in the configuration file.</p>
<div class="note" id="note891964620819"><a name="note891964620819"></a><a name="note891964620819"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p1091964618820"><a name="p1091964618820"></a><a name="p1091964618820"></a>The value is ensured to be greater than or equal to 1.</p>
</div></div>
</td>
</tr>
<tr id="row440921733714"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="24%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="60%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

>![](public_sys-resources/icon-note.gif) **NOTE:**   
>Only one of  **sc**,  **acl**, and  **aclXml**  can be set for each command.  

## Response<a name="section6926520122416"></a>

Refer to  [Response](uploading-an-object.md#section6926520122416)  for uploading an object.

## Running Examples<a name="section15899161919244"></a>

-   Take the Windows OS as an example, run the  **obsutil chattri obs://bucket-test/key -acl=public-read**  command to set the access permission to an object to public read.

```
obsutil chattri obs://bucket-test/key -acl=public-read

Set the acl of object [key] in the bucket [bucket-test] to [public-read] successfully, request id [04050000016836DDFA73B2B5320E2651]
```

-   Take the Windows OS as an example, run the  **obsutil  **chattri **obs://bucket-test -r -f  **-acl=public-read****  command to set the access permission to all objects in the bucket to public read.

```
obsutil chattri obs://bucket-test -r -f -acl=public-read

[------------------------------------------------] 100.00% tps:155.15 5/5 233ms
Succeed count is:   5         Failed count is:    0
Metrics [max cost:177 ms, min cost:53 ms, average cost:102.40 ms, average tps:20.41]
Task id is: 9d7f73ff-f747-4fdd-9b2a-815ba2dc3b07
```

