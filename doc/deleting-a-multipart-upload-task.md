# Deleting a Multipart Upload Task<a name="EN-US_TOPIC_0141181366"></a>

## Function<a name="section1479112110815"></a>

-   You can use this command to delete a multipart upload task in a specified bucket by using the multipart upload ID.
-   You can also use this command to delete multipart upload tasks in batches based on a specified object name prefix.

## Command Line Structure<a name="section1220752192216"></a>

-   In Windows
    -   Deleting a single multipart upload task

        ```
        obsutil abort obs://bucket/key -u=xxx [-f] [-fr] [-o=xxx] [-config=xxx]
        ```

    -   Deleting multipart upload tasks in batches

        ```
        obsutil abort obs://bucket[/key] -r [-f] [-o=xxx] [-j=1] [-config=xxx]
        ```


-   In Linux or macOS
    -   Deleting a single multipart upload task

        ```
        ./obsutil abort obs://bucket/key -u=xxx [-f] [-fr] [-o=xxx] [-config=xxx]
        ```

    -   Deleting multipart upload tasks in batches

        ```
        ./obsutil abort obs://bucket[/key] -r [-f] [-o=xxx] [-j=1] [-config=xxx]
        ```



## Parameter Description<a name="section6559191102418"></a>

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="17%" id="mcps1.1.4.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="28.000000000000004%" id="mcps1.1.4.1.2"><p id="p15137125919108"><a name="p15137125919108"></a><a name="p15137125919108"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="55.00000000000001%" id="mcps1.1.4.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row108328217449"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p64495172515"><a name="p64495172515"></a><a name="p64495172515"></a>bucket</p>
</td>
<td class="cellrowborder" valign="top" width="28.000000000000004%" headers="mcps1.1.4.1.2 "><p id="p154316502519"><a name="p154316502519"></a><a name="p154316502519"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p17425512259"><a name="p17425512259"></a><a name="p17425512259"></a>Bucket name</p>
</td>
</tr>
<tr id="row20647122314615"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p864719233464"><a name="p864719233464"></a><a name="p864719233464"></a>key</p>
</td>
<td class="cellrowborder" valign="top" width="28.000000000000004%" headers="mcps1.1.4.1.2 "><p id="p52017286567"><a name="p52017286567"></a><a name="p52017286567"></a>Mandatory for deleting a multipart upload task.</p>
<p id="p741812412475"><a name="p741812412475"></a><a name="p741812412475"></a>Optional for deleting multipart upload tasks in batches.</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1864718233461"><a name="p1864718233461"></a><a name="p1864718233461"></a>Indicates the object name involved in a multipart upload task to be deleted, or the name prefix of the objects involved in multipart upload tasks to be deleted in batches.</p>
<div class="note" id="note7180183222110"><a name="note7180183222110"></a><a name="note7180183222110"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p81802321215"><a name="p81802321215"></a><a name="p81802321215"></a>If this parameter is left blank when deleting multipart upload tasks in batches, all multipart upload tasks in the bucket are deleted.</p>
</div></div>
</td>
</tr>
<tr id="row15278410149"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p152884115145"><a name="p152884115145"></a><a name="p152884115145"></a>u</p>
</td>
<td class="cellrowborder" valign="top" width="28.000000000000004%" headers="mcps1.1.4.1.2 "><p id="p12874171412"><a name="p12874171412"></a><a name="p12874171412"></a>Mandatory for deleting a single multipart upload task (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p63113823"><a name="p63113823"></a><a name="p63113823"></a>ID of the multipart upload task to be deleted</p>
<div class="note" id="note1295362512190"><a name="note1295362512190"></a><a name="note1295362512190"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p2953162521914"><a name="p2953162521914"></a><a name="p2953162521914"></a>You can obtain the value of this parameter from <a href="listing-multipart-upload-tasks.md">Listing Multipart Upload Tasks</a>.</p>
</div></div>
</td>
</tr>
<tr id="row135170431481"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p8517143114816"><a name="p8517143114816"></a><a name="p8517143114816"></a>fr</p>
</td>
<td class="cellrowborder" valign="top" width="28.000000000000004%" headers="mcps1.1.4.1.2 "><p id="p1965455113483"><a name="p1965455113483"></a><a name="p1965455113483"></a>Optional for deleting a single multipart upload task (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1951718433487"><a name="p1951718433487"></a><a name="p1951718433487"></a>Generates an operation result list when deleting a multipart upload task.</p>
</td>
</tr>
<tr id="row536914065314"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p9369940145311"><a name="p9369940145311"></a><a name="p9369940145311"></a>f</p>
</td>
<td class="cellrowborder" valign="top" width="28.000000000000004%" headers="mcps1.1.4.1.2 "><p id="p1272453612483"><a name="p1272453612483"></a><a name="p1272453612483"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p972417362487"><a name="p972417362487"></a><a name="p972417362487"></a>Runs in force mode.</p>
</td>
</tr>
<tr id="row1215934213533"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p71594425537"><a name="p71594425537"></a><a name="p71594425537"></a>r</p>
</td>
<td class="cellrowborder" valign="top" width="28.000000000000004%" headers="mcps1.1.4.1.2 "><p id="p1273555018488"><a name="p1273555018488"></a><a name="p1273555018488"></a>Mandatory for deleting multipart upload tasks (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1173519509481"><a name="p1173519509481"></a><a name="p1173519509481"></a>Deletes multipart upload tasks in batches based on a specified object name prefix.</p>
</td>
</tr>
<tr id="row616420205485"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p187441222114815"><a name="p187441222114815"></a><a name="p187441222114815"></a>j</p>
</td>
<td class="cellrowborder" valign="top" width="28.000000000000004%" headers="mcps1.1.4.1.2 "><p id="p17745112212488"><a name="p17745112212488"></a><a name="p17745112212488"></a>Optional for deleting multipart upload tasks (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p47481922194811"><a name="p47481922194811"></a><a name="p47481922194811"></a>Indicates the maximum number of concurrent tasks for deleting multipart uploads in batches. The default value is the value of <strong id="b16444613907"><a name="b16444613907"></a><a name="b16444613907"></a>defaultJobs</strong> in the configuration file.</p>
<div class="note" id="note891964620819"><a name="note891964620819"></a><a name="note891964620819"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p1091964618820"><a name="p1091964618820"></a><a name="p1091964618820"></a>The value is ensured to be greater than or equal to 1.</p>
</div></div>
</td>
</tr>
<tr id="row11857124714532"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p138574472533"><a name="p138574472533"></a><a name="p138574472533"></a>o</p>
</td>
<td class="cellrowborder" valign="top" width="28.000000000000004%" headers="mcps1.1.4.1.2 "><p id="p7281628155011"><a name="p7281628155011"></a><a name="p7281628155011"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p4924117181112"><a name="p4924117181112"></a><a name="p4924117181112"></a>Indicates the folder where operation result lists reside. After the command is executed, result lists (possibly including success and failure files) are generated in the folder. The default value is <strong id="b464252132112"><a name="b464252132112"></a><a name="b464252132112"></a>.obsutil_output</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note351244972418"><a name="note351244972418"></a><a name="note351244972418"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul101190347408"></a><a name="ul101190347408"></a><ul id="ul101190347408"><li>The naming rule for result lists is as follows: <strong id="b1994120441411"><a name="b1994120441411"></a><a name="b1994120441411"></a>abort_{succeed  | failed}_report_</strong><em id="i109451441842"><a name="i109451441842"></a><a name="i109451441842"></a>time</em><strong id="b16948194419413"><a name="b16948194419413"></a><a name="b16948194419413"></a>_TaskId.txt</strong></li><li>By default, the maximum size of a single result list is 30 MB and the maximum number of result lists that can be retained is 1024. You can set the maximum size and number by configuring <strong id="b29141448392"><a name="b29141448392"></a><a name="b29141448392"></a>recordMaxLogSize</strong> and <strong id="b5915164443916"><a name="b5915164443916"></a><a name="b5915164443916"></a>recordBackups</strong> in the configuration file.</li></ul>
</div></div>
</td>
</tr>
<tr id="row81209083710"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="28.000000000000004%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Response<a name="section6926520122416"></a>

Refer to  [Response](uploading-an-object.md#section6926520122416)  for uploading an object.

## Running Examples<a name="section15899161919244"></a>

-   Take the Windows OS as an example. Run the  **obsutil abort obs://bucket-test/key -u=xxx -f**  command to delete a single multipart upload task.

```
obsutil abort obs://bucket-test/key -u=xxx -f

Abort multipart upload [key] in the bucket [bucket-test] successfully!
```

-   Take the Windows OS as an example. Run the  **obsutil abort obs://bucket-test -r -f**  command to delete all multipart upload tasks in the bucket in batches.

```
obsutil abort obs://bucket-test -r -f

Listing multipart uploads to abort...
Aborting progress:[================================================] 100.00% 3s
Succeed count is:   12        Failed count is:    0
Metrics [max cost:264 ms, min cost:54 ms, average cost:119.33 ms, average tps:19.70]
Task id is: 0b34b1fa-b015-4313-a216-0fd5b4fffa1c
```

