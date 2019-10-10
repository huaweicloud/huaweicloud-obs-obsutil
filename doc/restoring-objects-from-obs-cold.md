# Restoring Objects from OBS Cold<a name="EN-US_TOPIC_0141181367"></a>

## Function<a name="section1479112110815"></a>

You can use this command to restore a specified object whose storage class is  **cold**  or restore objects in batches by object name prefix.

>![](public_sys-resources/icon-note.gif) **NOTE:**   
>-   Object content cannot be read during restoration.  
>-   After an object is restored, the time it requires before the object can be downloaded depends on the OBS server.  

## Command Line Structure<a name="section1220752192216"></a>

-   In Windows
    -   Restoring an object

        ```
        obsutil restore obs://bucket/key [-d=1] [-t=xxx] [-versionId=xxx] [-fr] [-o=xxx] [-config=xxx]
        ```

    -   Restoring objects in batches

        ```
        obsutil restore obs://bucket[/key] -r [-f] [-v] [-d=1] [-t=xxx] [-o=xxx] [-j=1] [-config=xxx]
        ```


-   In Linux or macOS
    -   Restoring an object

        ```
        ./obsutil restore obs://bucket/key [-d=1] [-t=xxx] [-versionId=xxx] [-fr] [-o=xxx] [-config=xxx]
        ```

    -   Restoring objects in batches

        ```
        ./obsutil restore obs://bucket[/key] -r [-f] [-v] [-d=1] [-t=xxx] [-o=xxx] [-j=1] [-config=xxx]
        ```



## Parameter Description<a name="section6559191102418"></a>

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="16.84%" id="mcps1.1.4.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="24.7%" id="mcps1.1.4.1.2"><p id="p15137125919108"><a name="p15137125919108"></a><a name="p15137125919108"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="58.46%" id="mcps1.1.4.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row108328217449"><td class="cellrowborder" valign="top" width="16.84%" headers="mcps1.1.4.1.1 "><p id="p64495172515"><a name="p64495172515"></a><a name="p64495172515"></a>bucket</p>
</td>
<td class="cellrowborder" valign="top" width="24.7%" headers="mcps1.1.4.1.2 "><p id="p154316502519"><a name="p154316502519"></a><a name="p154316502519"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="58.46%" headers="mcps1.1.4.1.3 "><p id="p17425512259"><a name="p17425512259"></a><a name="p17425512259"></a>Bucket name</p>
</td>
</tr>
<tr id="row14581162223916"><td class="cellrowborder" valign="top" width="16.84%" headers="mcps1.1.4.1.1 "><p id="p864719233464"><a name="p864719233464"></a><a name="p864719233464"></a>key</p>
</td>
<td class="cellrowborder" valign="top" width="24.7%" headers="mcps1.1.4.1.2 "><p id="p8166018195614"><a name="p8166018195614"></a><a name="p8166018195614"></a>Mandatory for restoring a single object whose storage class is <strong id="b229419472558"><a name="b229419472558"></a><a name="b229419472558"></a>cold</strong></p>
<p id="p741812412475"><a name="p741812412475"></a><a name="p741812412475"></a>Optional for batch restoring objects whose storage class is <strong id="b1943244175720"><a name="b1943244175720"></a><a name="b1943244175720"></a>cold</strong></p>
</td>
<td class="cellrowborder" valign="top" width="58.46%" headers="mcps1.1.4.1.3 "><p id="p1864718233461"><a name="p1864718233461"></a><a name="p1864718233461"></a>Indicates the name of the object to be restored or the name prefix of the objects to be restored in batches.</p>
<div class="note" id="note12611558164215"><a name="note12611558164215"></a><a name="note12611558164215"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p1061114584426"><a name="p1061114584426"></a><a name="p1061114584426"></a>If this parameter is left blank when batch restoring objects, all objects whose storage class is <strong id="b14694184835810"><a name="b14694184835810"></a><a name="b14694184835810"></a>cold</strong> in the bucket are restored.</p>
</div></div>
</td>
</tr>
<tr id="row5136104314396"><td class="cellrowborder" valign="top" width="16.84%" headers="mcps1.1.4.1.1 "><p id="p152884115145"><a name="p152884115145"></a><a name="p152884115145"></a>d</p>
</td>
<td class="cellrowborder" valign="top" width="24.7%" headers="mcps1.1.4.1.2 "><p id="p12874171412"><a name="p12874171412"></a><a name="p12874171412"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.46%" headers="mcps1.1.4.1.3 "><p id="p63113823"><a name="p63113823"></a><a name="p63113823"></a>Storage duration after objects whose storage class is <strong id="b33817465411"><a name="b33817465411"></a><a name="b33817465411"></a>cold</strong> are restored, in days. The value ranges from 1 to 30. The default value is 1.</p>
</td>
</tr>
<tr id="row1979619354115"><td class="cellrowborder" valign="top" width="16.84%" headers="mcps1.1.4.1.1 "><p id="p187966384115"><a name="p187966384115"></a><a name="p187966384115"></a>t</p>
</td>
<td class="cellrowborder" valign="top" width="24.7%" headers="mcps1.1.4.1.2 "><p id="p97961933415"><a name="p97961933415"></a><a name="p97961933415"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.46%" headers="mcps1.1.4.1.3 "><p id="p579611319412"><a name="p579611319412"></a><a name="p579611319412"></a>Options for restoring objects. Possible values are:</p>
<a name="ul1273864144412"></a><a name="ul1273864144412"></a><ul id="ul1273864144412"><li>standard</li><li>expedited</li></ul>
<div class="note" id="note1980113505431"><a name="note1980113505431"></a><a name="note1980113505431"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul6317738122314"></a><a name="ul6317738122314"></a><ul id="ul6317738122314"><li>The preceding two values indicate standard restoration (3–5 hours) and quick restoration (1–5 minutes).</li><li>If this parameter is not set, expedited restoration is used by default.</li></ul>
</div></div>
</td>
</tr>
<tr id="row1742147104113"><td class="cellrowborder" valign="top" width="16.84%" headers="mcps1.1.4.1.1 "><p id="p174267194117"><a name="p174267194117"></a><a name="p174267194117"></a>versionId</p>
</td>
<td class="cellrowborder" valign="top" width="24.7%" headers="mcps1.1.4.1.2 "><p id="p774216710415"><a name="p774216710415"></a><a name="p774216710415"></a>Optional for restoring a single object whose storage class is <strong id="b194658220715"><a name="b194658220715"></a><a name="b194658220715"></a>cold</strong> (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.46%" headers="mcps1.1.4.1.3 "><p id="p874257164113"><a name="p874257164113"></a><a name="p874257164113"></a>Version ID of the to-be-restored object whose storage class is <strong id="b742771211115"><a name="b742771211115"></a><a name="b742771211115"></a>cold</strong></p>
</td>
</tr>
<tr id="row1671701015561"><td class="cellrowborder" valign="top" width="16.84%" headers="mcps1.1.4.1.1 "><p id="p8517143114816"><a name="p8517143114816"></a><a name="p8517143114816"></a>fr</p>
</td>
<td class="cellrowborder" valign="top" width="24.7%" headers="mcps1.1.4.1.2 "><p id="p1965455113483"><a name="p1965455113483"></a><a name="p1965455113483"></a>Optional for restoring a single object whose storage class is <strong id="b519211231219"><a name="b519211231219"></a><a name="b519211231219"></a>cold</strong> (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.46%" headers="mcps1.1.4.1.3 "><p id="p1951718433487"><a name="p1951718433487"></a><a name="p1951718433487"></a>Generates an operation result list when restoring a single object whose storage class is <strong id="b1932442821220"><a name="b1932442821220"></a><a name="b1932442821220"></a>cold</strong>.</p>
</td>
</tr>
<tr id="row18931643133919"><td class="cellrowborder" valign="top" width="16.84%" headers="mcps1.1.4.1.1 "><p id="p9369940145311"><a name="p9369940145311"></a><a name="p9369940145311"></a>f</p>
</td>
<td class="cellrowborder" valign="top" width="24.7%" headers="mcps1.1.4.1.2 "><p id="p19278143834816"><a name="p19278143834816"></a><a name="p19278143834816"></a>Optional for batch restoring objects whose storage class is <strong id="b850354517132"><a name="b850354517132"></a><a name="b850354517132"></a>cold</strong> (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.46%" headers="mcps1.1.4.1.3 "><p id="p972417362487"><a name="p972417362487"></a><a name="p972417362487"></a>Runs in force mode.</p>
</td>
</tr>
<tr id="row11636444193916"><td class="cellrowborder" valign="top" width="16.84%" headers="mcps1.1.4.1.1 "><p id="p71594425537"><a name="p71594425537"></a><a name="p71594425537"></a>r</p>
</td>
<td class="cellrowborder" valign="top" width="24.7%" headers="mcps1.1.4.1.2 "><p id="p1273555018488"><a name="p1273555018488"></a><a name="p1273555018488"></a>Mandatory for batch restoring objects whose storage class is <strong id="b3389249149"><a name="b3389249149"></a><a name="b3389249149"></a>cold</strong> (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.46%" headers="mcps1.1.4.1.3 "><p id="p1173519509481"><a name="p1173519509481"></a><a name="p1173519509481"></a>Restores objects whose storage class is <strong id="b3789244131412"><a name="b3789244131412"></a><a name="b3789244131412"></a>cold</strong> in batches by object name prefix.</p>
</td>
</tr>
<tr id="row07312318419"><td class="cellrowborder" valign="top" width="16.84%" headers="mcps1.1.4.1.1 "><p id="p373223134117"><a name="p373223134117"></a><a name="p373223134117"></a>v</p>
</td>
<td class="cellrowborder" valign="top" width="24.7%" headers="mcps1.1.4.1.2 "><p id="p101701338485"><a name="p101701338485"></a><a name="p101701338485"></a>Optional for batch restoring objects whose storage class is <strong id="b9428415141618"><a name="b9428415141618"></a><a name="b9428415141618"></a>cold</strong> (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.46%" headers="mcps1.1.4.1.3 "><p id="p1073202316412"><a name="p1073202316412"></a><a name="p1073202316412"></a>Restores versions of objects whose storage class is <strong id="b217192141719"><a name="b217192141719"></a><a name="b217192141719"></a>cold</strong> in batches by object name prefix.</p>
</td>
</tr>
<tr id="row102561945153913"><td class="cellrowborder" valign="top" width="16.84%" headers="mcps1.1.4.1.1 "><p id="p138574472533"><a name="p138574472533"></a><a name="p138574472533"></a>o</p>
</td>
<td class="cellrowborder" valign="top" width="24.7%" headers="mcps1.1.4.1.2 "><p id="p817213354812"><a name="p817213354812"></a><a name="p817213354812"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.46%" headers="mcps1.1.4.1.3 "><p id="p4924117181112"><a name="p4924117181112"></a><a name="p4924117181112"></a>Indicates the folder where operation result lists reside. After the command is executed, result lists (possibly including success and failure files) are generated in the folder. The default value is <strong id="b317943718613"><a name="b317943718613"></a><a name="b317943718613"></a>.obsutil_output</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note349481802414"><a name="note349481802414"></a><a name="note349481802414"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul101190347408"></a><a name="ul101190347408"></a><ul id="ul101190347408"><li>The naming rule for result lists is as follows: <strong id="b279801517336"><a name="b279801517336"></a><a name="b279801517336"></a>restore_{succeed  | failed}_report_</strong><em id="i079981513335"><a name="i079981513335"></a><a name="i079981513335"></a>time</em><strong id="b10799915143310"><a name="b10799915143310"></a><a name="b10799915143310"></a>_TaskId.txt</strong><p id="li412103664016p0"><a name="li412103664016p0"></a><a name="li412103664016p0"></a>By default, the maximum size of a single result list is 30 MB and the maximum number of result lists that can be retained is 1024. You can set the maximum size and number by configuring <strong id="b15561958133811"><a name="b15561958133811"></a><a name="b15561958133811"></a>recordMaxLogSize</strong> and <strong id="b05613585389"><a name="b05613585389"></a><a name="b05613585389"></a>recordBackups</strong> in the configuration file.</p>
</li></ul>
</div></div>
</td>
</tr>
<tr id="row588544593912"><td class="cellrowborder" valign="top" width="16.84%" headers="mcps1.1.4.1.1 "><p id="p87050508534"><a name="p87050508534"></a><a name="p87050508534"></a>j</p>
</td>
<td class="cellrowborder" valign="top" width="24.7%" headers="mcps1.1.4.1.2 "><p id="p19177173320483"><a name="p19177173320483"></a><a name="p19177173320483"></a>Optional for batch restoring objects whose storage class is <strong id="b488113297181"><a name="b488113297181"></a><a name="b488113297181"></a>cold</strong> (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.46%" headers="mcps1.1.4.1.3 "><p id="p2549177181312"><a name="p2549177181312"></a><a name="p2549177181312"></a>Maximum number of concurrent tasks for batch restoring objects whose storage class is <strong id="b10998114415196"><a name="b10998114415196"></a><a name="b10998114415196"></a>cold</strong>. The default value is the value of <strong id="b1080653133415"><a name="b1080653133415"></a><a name="b1080653133415"></a>defaultJobs</strong> in the configuration file.</p>
<div class="note" id="note891964620819"><a name="note891964620819"></a><a name="note891964620819"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p1091964618820"><a name="p1091964618820"></a><a name="p1091964618820"></a>The value is ensured to be greater than or equal to 1.</p>
</div></div>
</td>
</tr>
<tr id="row55581242153513"><td class="cellrowborder" valign="top" width="16.84%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="24.7%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="58.46%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Response<a name="section6926520122416"></a>

Refer to  [Response](uploading-an-object.md#section6926520122416)  for uploading an object.

## Running Examples<a name="section15899161919244"></a>

-   Take the Windows OS as an example. Run the  **obsutil restore obs://bucket-test/key**  command to restore a single object whose storage class is  **cold**.

```
obsutil restore obs://bucket-test/key

Start to restore object [key] in the bucket [bucket-test] successfully!
```

-   Take the Windows OS as an example. Run the  **obsutil restore obs://bucket-test -r -f**  command to restore objects whose storage class is  **cold**  in the bucket in batches.

```
obsutil restore obs://bucket-test -r -f

[================================================] 100.00% 3s
Succeed count is:   12        Failed count is:    0
Metrics [max cost:264 ms, min cost:54 ms, average cost:119.33 ms, average tps:19.70]
Task id is: 96f104ee-d0bf-40ff-95dd-31dec0d8f4f4
```

