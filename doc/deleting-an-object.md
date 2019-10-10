# Deleting an Object<a name="EN-US_TOPIC_0142009731"></a>

## Function<a name="section1479112110815"></a>

-   You can use this command to delete a specified object.
-   You can also use this command to delete objects in batches based on a specified object name prefix.

## Command Line Structure<a name="section1220752192216"></a>

-   In Windows
    -   Deleting a single object

        ```
        obsutil rm obs://bucket/key [-f] [-versionId=xxx] [-fr] [-o=xxx] [-config=xxx]
        ```

    -   Deleting objects in batches

        ```
        obsutil rm obs://bucket/[key] -r [-j=1] [-f] [-v] [-o=xxx] [-config=xxx]
        ```


-   In Linux or macOS
    -   Deleting a single object

        ```
        ./obsutil rm obs://bucket/key [-f] [-versionId=xxx] [-fr] [-o=xxx] [-config=xxx]
        ```

    -   Deleting objects in batches

        ```
        ./obsutil rm obs://bucket/[key] -r [-j=1] [-f] [-v] [-o=xxx] [-config=xxx]
        ```



## Parameter Description<a name="section6559191102418"></a>

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="16%" id="mcps1.1.4.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="26%" id="mcps1.1.4.1.2"><p id="p15137125919108"><a name="p15137125919108"></a><a name="p15137125919108"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="57.99999999999999%" id="mcps1.1.4.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row108328217449"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p64495172515"><a name="p64495172515"></a><a name="p64495172515"></a>bucket</p>
</td>
<td class="cellrowborder" valign="top" width="26%" headers="mcps1.1.4.1.2 "><p id="p154316502519"><a name="p154316502519"></a><a name="p154316502519"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="57.99999999999999%" headers="mcps1.1.4.1.3 "><p id="p17425512259"><a name="p17425512259"></a><a name="p17425512259"></a>Bucket name</p>
</td>
</tr>
<tr id="row14181641114718"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p2418114124718"><a name="p2418114124718"></a><a name="p2418114124718"></a>key</p>
</td>
<td class="cellrowborder" valign="top" width="26%" headers="mcps1.1.4.1.2 "><p id="p76483210569"><a name="p76483210569"></a><a name="p76483210569"></a>Mandatory for deleting a single object.</p>
<p id="p741812412475"><a name="p741812412475"></a><a name="p741812412475"></a>Optional for deleting objects in batches.</p>
</td>
<td class="cellrowborder" valign="top" width="57.99999999999999%" headers="mcps1.1.4.1.3 "><p id="p741894118475"><a name="p741894118475"></a><a name="p741894118475"></a>Indicates the name of the object to be deleted, or the name prefix of the objects to be deleted in batches.</p>
<div class="note" id="note7180183222110"><a name="note7180183222110"></a><a name="note7180183222110"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p81802321215"><a name="p81802321215"></a><a name="p81802321215"></a>If this parameter is left blank when deleting objects in batches, all objects in the bucket are deleted.</p>
</div></div>
</td>
</tr>
<tr id="row20170121265812"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p8517143114816"><a name="p8517143114816"></a><a name="p8517143114816"></a>fr</p>
</td>
<td class="cellrowborder" valign="top" width="26%" headers="mcps1.1.4.1.2 "><p id="p1965455113483"><a name="p1965455113483"></a><a name="p1965455113483"></a>Optional for deleting a single object (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="57.99999999999999%" headers="mcps1.1.4.1.3 "><p id="p1951718433487"><a name="p1951718433487"></a><a name="p1951718433487"></a>Generates an operation result list when deleting an object.</p>
</td>
</tr>
<tr id="row107241436114819"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p5724123684813"><a name="p5724123684813"></a><a name="p5724123684813"></a>f</p>
</td>
<td class="cellrowborder" valign="top" width="26%" headers="mcps1.1.4.1.2 "><p id="p1272453612483"><a name="p1272453612483"></a><a name="p1272453612483"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="57.99999999999999%" headers="mcps1.1.4.1.3 "><p id="p972417362487"><a name="p972417362487"></a><a name="p972417362487"></a>Runs in force mode.</p>
</td>
</tr>
<tr id="row1460034216484"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p560014284815"><a name="p560014284815"></a><a name="p560014284815"></a>versionId</p>
</td>
<td class="cellrowborder" valign="top" width="26%" headers="mcps1.1.4.1.2 "><p id="p10600184264818"><a name="p10600184264818"></a><a name="p10600184264818"></a>Optional for deleting a single object (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="57.99999999999999%" headers="mcps1.1.4.1.3 "><p id="p56001942134812"><a name="p56001942134812"></a><a name="p56001942134812"></a>Version ID of the object to be deleted.</p>
</td>
</tr>
<tr id="row18735850114810"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p3735185044816"><a name="p3735185044816"></a><a name="p3735185044816"></a>r</p>
</td>
<td class="cellrowborder" valign="top" width="26%" headers="mcps1.1.4.1.2 "><p id="p15786145522217"><a name="p15786145522217"></a><a name="p15786145522217"></a>Mandatory for deleting objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="57.99999999999999%" headers="mcps1.1.4.1.3 "><p id="p1173519509481"><a name="p1173519509481"></a><a name="p1173519509481"></a>Deletes objects in batches based on a specified object name prefix.</p>
</td>
</tr>
<tr id="row015935193510"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p3662204585313"><a name="p3662204585313"></a><a name="p3662204585313"></a>j</p>
</td>
<td class="cellrowborder" valign="top" width="26%" headers="mcps1.1.4.1.2 "><p id="p19305175612615"><a name="p19305175612615"></a><a name="p19305175612615"></a>Optional for deleting objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="57.99999999999999%" headers="mcps1.1.4.1.3 "><p id="p17665114555311"><a name="p17665114555311"></a><a name="p17665114555311"></a>Indicates the maximum number of concurrent tasks for deleting objects in batches. The default value is the value of <strong id="b1268754012466"><a name="b1268754012466"></a><a name="b1268754012466"></a>defaultJobs</strong> in the configuration file.</p>
<div class="note" id="note891964620819"><a name="note891964620819"></a><a name="note891964620819"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p1091964618820"><a name="p1091964618820"></a><a name="p1091964618820"></a>The value is ensured to be greater than or equal to 1.</p>
</div></div>
</td>
</tr>
<tr id="row198551816104916"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p20856121619492"><a name="p20856121619492"></a><a name="p20856121619492"></a>v</p>
</td>
<td class="cellrowborder" valign="top" width="26%" headers="mcps1.1.4.1.2 "><p id="p3791455132212"><a name="p3791455132212"></a><a name="p3791455132212"></a>Optional for deleting objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="57.99999999999999%" headers="mcps1.1.4.1.3 "><p id="p1385671610497"><a name="p1385671610497"></a><a name="p1385671610497"></a>Deletes versions of an object and the delete markers in batches based on a specified object name prefix.</p>
</td>
</tr>
<tr id="row106825331499"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p268213314499"><a name="p268213314499"></a><a name="p268213314499"></a>o</p>
</td>
<td class="cellrowborder" valign="top" width="26%" headers="mcps1.1.4.1.2 "><p id="p1979414554228"><a name="p1979414554228"></a><a name="p1979414554228"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="57.99999999999999%" headers="mcps1.1.4.1.3 "><p id="p10682123310498"><a name="p10682123310498"></a><a name="p10682123310498"></a>Indicates the folder where operation result lists reside. After the command is executed, result lists (possibly including success and failure files) are generated in the folder. The default value is <strong id="b2095215720228"><a name="b2095215720228"></a><a name="b2095215720228"></a>.obsutil_output</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note236970182510"><a name="note236970182510"></a><a name="note236970182510"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul101190347408"></a><a name="ul101190347408"></a><ul id="ul101190347408"><li>The naming rule for result lists is as follows: <strong id="b9923175232616"><a name="b9923175232616"></a><a name="b9923175232616"></a>rm_{succeed  | failed}_report_</strong><em id="i17925205219267"><a name="i17925205219267"></a><a name="i17925205219267"></a>time</em><strong id="b392511527263"><a name="b392511527263"></a><a name="b392511527263"></a>_TaskId.txt</strong></li><li>By default, the maximum size of a single result list is 30 MB and the maximum number of result lists that can be retained is 1024. You can set the maximum size and number by configuring <strong id="b357011503361"><a name="b357011503361"></a><a name="b357011503361"></a>recordMaxLogSize</strong> and <strong id="b157095015367"><a name="b157095015367"></a><a name="b157095015367"></a>recordBackups</strong> in the configuration file.</li></ul>
</div></div>
</td>
</tr>
<tr id="row083396143714"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="26%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="57.99999999999999%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Response<a name="section6926520122416"></a>

Refer to  [Response](uploading-an-object.md#section6926520122416)  for uploading an object.

## Running Examples<a name="section15899161919244"></a>

-   Take the Windows OS as an example. Run the  **obsutil rm obs://bucket-test/key -f**  command to delete a single object named  **key**  in bucket  **bucket-test**.

```
obsutil rm obs://bucket-test/key -f

Delete object [key] in the bucket [bucket-test] successfully!
```

-   Take the Windows OS as an example. Run the  **obsutil rm obs://bucket-test -r -f**  command to delete all objects in bucket  **bucket-test**.

```
obsutil rm obs://bucket-test -r -f

[===============================================] 100.00% 21s
Succeed count is:   1313      Failed count is:    0
Task id is: 95936984-f81a-441a-bba0-1fd8254d9241
```

