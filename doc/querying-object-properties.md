# Querying Object Properties<a name="EN-US_TOPIC_0142009730"></a>

## Function<a name="section444185010363"></a>

You can use this command to query the basic properties of an object.

## Command Line Structure<a name="section1220752192216"></a>

-   In Windows

    ```
    obsutil stat obs://bucket/key [-acl] [-config=xxx]
    ```

-   In Linux or macOS

    ```
    ./obsutil stat obs://bucket/key [-acl] [-config=xxx]
    ```


## Parameter Description<a name="section6559191102418"></a>

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="17%" id="mcps1.1.4.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="27%" id="mcps1.1.4.1.2"><p id="p15137125919108"><a name="p15137125919108"></a><a name="p15137125919108"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="56.00000000000001%" id="mcps1.1.4.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row108328217449"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p64495172515"><a name="p64495172515"></a><a name="p64495172515"></a>bucket</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p154316502519"><a name="p154316502519"></a><a name="p154316502519"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="56.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p17425512259"><a name="p17425512259"></a><a name="p17425512259"></a>Bucket name</p>
</td>
</tr>
<tr id="row178507135421"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p18851121313427"><a name="p18851121313427"></a><a name="p18851121313427"></a>key</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p1285119130421"><a name="p1285119130421"></a><a name="p1285119130421"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="56.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1685111317426"><a name="p1685111317426"></a><a name="p1685111317426"></a>Object name</p>
</td>
</tr>
<tr id="row1542113212124"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p814111533115"><a name="p814111533115"></a><a name="p814111533115"></a>acl</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p181412534110"><a name="p181412534110"></a><a name="p181412534110"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="56.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p20141165311112"><a name="p20141165311112"></a><a name="p20141165311112"></a>Queries the access control policies of the object at the same time.</p>
</td>
</tr>
<tr id="row17576193493515"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Response<a name="section6926520122416"></a>

<a name="table992610203244"></a>
<table><thead align="left"><tr id="row892913208248"><th class="cellrowborder" valign="top" width="21.67%" id="mcps1.1.3.1.1"><p id="p1392992019245"><a name="p1392992019245"></a><a name="p1392992019245"></a>Field</p>
</th>
<th class="cellrowborder" valign="top" width="78.33%" id="mcps1.1.3.1.2"><p id="p19318207246"><a name="p19318207246"></a><a name="p19318207246"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row79320208248"><td class="cellrowborder" valign="top" width="21.67%" headers="mcps1.1.3.1.1 "><p id="p1093217203246"><a name="p1093217203246"></a><a name="p1093217203246"></a>Key</p>
</td>
<td class="cellrowborder" valign="top" width="78.33%" headers="mcps1.1.3.1.2 "><p id="p169337207245"><a name="p169337207245"></a><a name="p169337207245"></a>Object name</p>
</td>
</tr>
<tr id="row1753615812247"><td class="cellrowborder" valign="top" width="21.67%" headers="mcps1.1.3.1.1 "><p id="p15536155820245"><a name="p15536155820245"></a><a name="p15536155820245"></a>LastModified</p>
</td>
<td class="cellrowborder" valign="top" width="78.33%" headers="mcps1.1.3.1.2 "><p id="p3536758162410"><a name="p3536758162410"></a><a name="p3536758162410"></a>Last modification time of the object</p>
</td>
</tr>
<tr id="row1272082862513"><td class="cellrowborder" valign="top" width="21.67%" headers="mcps1.1.3.1.1 "><p id="p4720528142516"><a name="p4720528142516"></a><a name="p4720528142516"></a>Size</p>
</td>
<td class="cellrowborder" valign="top" width="78.33%" headers="mcps1.1.3.1.2 "><p id="p872012286252"><a name="p872012286252"></a><a name="p872012286252"></a>Object size, in bytes</p>
</td>
</tr>
<tr id="row11100149145919"><td class="cellrowborder" valign="top" width="21.67%" headers="mcps1.1.3.1.1 "><p id="p210114914596"><a name="p210114914596"></a><a name="p210114914596"></a>StorageClass</p>
</td>
<td class="cellrowborder" valign="top" width="78.33%" headers="mcps1.1.3.1.2 "><p id="p11011093596"><a name="p11011093596"></a><a name="p11011093596"></a>Storage class of the object</p>
</td>
</tr>
<tr id="row10250185112481"><td class="cellrowborder" valign="top" width="21.67%" headers="mcps1.1.3.1.1 "><p id="p1525135124814"><a name="p1525135124814"></a><a name="p1525135124814"></a>MD5</p>
</td>
<td class="cellrowborder" valign="top" width="78.33%" headers="mcps1.1.3.1.2 "><p id="p930919517495"><a name="p930919517495"></a><a name="p930919517495"></a>Real MD5 of the object</p>
<div class="note" id="note1790518124911"><a name="note1790518124911"></a><a name="note1790518124911"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p8906128194916"><a name="p8906128194916"></a><a name="p8906128194916"></a>You can query this value only after running the <strong id="b648315819170"><a name="b648315819170"></a><a name="b648315819170"></a>cp</strong> command and configuring the <strong id="b14956204313177"><a name="b14956204313177"></a><a name="b14956204313177"></a>-vmd5</strong> parameter.</p>
</div></div>
</td>
</tr>
<tr id="row48453315255"><td class="cellrowborder" valign="top" width="21.67%" headers="mcps1.1.3.1.1 "><p id="p14845531172510"><a name="p14845531172510"></a><a name="p14845531172510"></a>ETag</p>
</td>
<td class="cellrowborder" valign="top" width="78.33%" headers="mcps1.1.3.1.2 "><p id="p11845131182516"><a name="p11845131182516"></a><a name="p11845131182516"></a>ETag value of an object calculated on the server</p>
</td>
</tr>
<tr id="row714915616589"><td class="cellrowborder" valign="top" width="21.67%" headers="mcps1.1.3.1.1 "><p id="p141743597587"><a name="p141743597587"></a><a name="p141743597587"></a>ContentType</p>
</td>
<td class="cellrowborder" valign="top" width="78.33%" headers="mcps1.1.3.1.2 "><p id="p10174175975817"><a name="p10174175975817"></a><a name="p10174175975817"></a>Content-Type of the object</p>
</td>
</tr>
<tr id="row9438135064914"><td class="cellrowborder" valign="top" width="21.67%" headers="mcps1.1.3.1.1 "><p id="p14438550144911"><a name="p14438550144911"></a><a name="p14438550144911"></a>Metadata</p>
</td>
<td class="cellrowborder" valign="top" width="78.33%" headers="mcps1.1.3.1.2 "><p id="p184387501497"><a name="p184387501497"></a><a name="p184387501497"></a>Customized metadata of the object</p>
</td>
</tr>
<tr id="row882474212589"><td class="cellrowborder" valign="top" width="21.67%" headers="mcps1.1.3.1.1 "><p id="p5286146135817"><a name="p5286146135817"></a><a name="p5286146135817"></a>Acl</p>
</td>
<td class="cellrowborder" valign="top" width="78.33%" headers="mcps1.1.3.1.2 "><p id="p182864618582"><a name="p182864618582"></a><a name="p182864618582"></a>Access control policy of the object</p>
</td>
</tr>
</tbody>
</table>

## Running Example<a name="section15899161919244"></a>

-   Take the Windows OS as an example. Run the  **obsutil stat obs://bucket-test/key**  command to query the basic properties of an object.

```
obsutil stat obs://bucket-test/key

Key:
  obs://bucket-test/key
LastModified:
  2018-11-16T02:15:49Z
Size:
  7
ETag:
  43d93b553855b0e1fc67e31c28c07b65
ContentType:
  text/plain
```

