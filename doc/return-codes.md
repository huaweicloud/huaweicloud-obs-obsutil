# Return Codes<a name="EN-US_TOPIC_0164918060"></a>

If obsutil is invoked by processes, the command output cannot be viewed in real time. obsutil generates different return codes based on different execution results.  [Table 1](#table936012247195)  describes the return codes. You can use either the following methods to obtain the return code of the latest execution result and then analyze and rectify the fault based on it:

-   In the macOS or Linux OS, run the following command to obtain the return code of the latest execution result:

    ```
    echo $?
    ```

-   In the Windows OS, run the following command to obtain the return code of the latest execution result:

    ```
    echo %errorlevel%
    ```


**Table  1**  Return codes

<a name="table936012247195"></a>
<table><thead align="left"><tr id="row13361112412193"><th class="cellrowborder" valign="top" width="18.09180918091809%" id="mcps1.2.4.1.1"><p id="p10361724141920"><a name="p10361724141920"></a><a name="p10361724141920"></a>Return Code</p>
</th>
<th class="cellrowborder" valign="top" width="26.82268226822682%" id="mcps1.2.4.1.2"><p id="p1836110246197"><a name="p1836110246197"></a><a name="p1836110246197"></a>Meaning</p>
</th>
<th class="cellrowborder" valign="top" width="55.08550855085508%" id="mcps1.2.4.1.3"><p id="p5361172414190"><a name="p5361172414190"></a><a name="p5361172414190"></a>Example Scenario</p>
</th>
</tr>
</thead>
<tbody><tr id="row15361324111912"><td class="cellrowborder" valign="top" width="18.09180918091809%" headers="mcps1.2.4.1.1 "><p id="p136132419192"><a name="p136132419192"></a><a name="p136132419192"></a>0</p>
</td>
<td class="cellrowborder" valign="top" width="26.82268226822682%" headers="mcps1.2.4.1.2 "><p id="p1736152401911"><a name="p1736152401911"></a><a name="p1736152401911"></a>Execution succeeded.</p>
</td>
<td class="cellrowborder" valign="top" width="55.08550855085508%" headers="mcps1.2.4.1.3 "><p id="p4361182421920"><a name="p4361182421920"></a><a name="p4361182421920"></a>An object is successfully uploaded.</p>
</td>
</tr>
<tr id="row83611124201915"><td class="cellrowborder" valign="top" width="18.09180918091809%" headers="mcps1.2.4.1.1 "><p id="p936172471916"><a name="p936172471916"></a><a name="p936172471916"></a>1</p>
</td>
<td class="cellrowborder" valign="top" width="26.82268226822682%" headers="mcps1.2.4.1.2 "><p id="p336172431919"><a name="p336172431919"></a><a name="p336172431919"></a>The file does not exist.</p>
</td>
<td class="cellrowborder" valign="top" width="55.08550855085508%" headers="mcps1.2.4.1.3 "><p id="p7361172415199"><a name="p7361172415199"></a><a name="p7361172415199"></a>The entered file path does not exist for uploading a file by running the <strong id="b7139122016518"><a name="b7139122016518"></a><a name="b7139122016518"></a>cp</strong> command.</p>
</td>
</tr>
<tr id="row425016264216"><td class="cellrowborder" valign="top" width="18.09180918091809%" headers="mcps1.2.4.1.1 "><p id="p825111269214"><a name="p825111269214"></a><a name="p825111269214"></a>2</p>
</td>
<td class="cellrowborder" valign="top" width="26.82268226822682%" headers="mcps1.2.4.1.2 "><p id="p7251626172119"><a name="p7251626172119"></a><a name="p7251626172119"></a>The task does not exist.</p>
</td>
<td class="cellrowborder" valign="top" width="55.08550855085508%" headers="mcps1.2.4.1.3 "><p id="p13251726132113"><a name="p13251726132113"></a><a name="p13251726132113"></a>The specified task ID does not exist for resuming a failed upload task by running the <strong id="b126989471670"><a name="b126989471670"></a><a name="b126989471670"></a>cp</strong> command.</p>
</td>
</tr>
<tr id="row15969153413237"><td class="cellrowborder" valign="top" width="18.09180918091809%" headers="mcps1.2.4.1.1 "><p id="p79691234122314"><a name="p79691234122314"></a><a name="p79691234122314"></a>3</p>
</td>
<td class="cellrowborder" valign="top" width="26.82268226822682%" headers="mcps1.2.4.1.2 "><p id="p6969734182317"><a name="p6969734182317"></a><a name="p6969734182317"></a>Parameter error</p>
</td>
<td class="cellrowborder" valign="top" width="55.08550855085508%" headers="mcps1.2.4.1.3 "><a name="ul169195316318"></a><a name="ul169195316318"></a><ul id="ul169195316318"><li>At least one entered additional parameters is not supported for uploading a file by running the <strong id="b1327518712111"><a name="b1327518712111"></a><a name="b1327518712111"></a>cp</strong> command.</li><li>The entered value of <strong id="b17316740171412"><a name="b17316740171412"></a><a name="b17316740171412"></a>cloud_url</strong> is invalid for downloading a file by running the <strong id="b36799509142"><a name="b36799509142"></a><a name="b36799509142"></a>cp</strong> command.<div class="note" id="note2064653215122"><a name="note2064653215122"></a><a name="note2064653215122"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p10647432201214"><a name="p10647432201214"></a><a name="p10647432201214"></a><strong id="b1256924015459"><a name="b1256924015459"></a><a name="b1256924015459"></a>cloud_url</strong> indicates the bucket path or object path. Set <strong id="b3717138105313"><a name="b3717138105313"></a><a name="b3717138105313"></a>cloud_url</strong> in the format of <strong id="b1680232913532"><a name="b1680232913532"></a><a name="b1680232913532"></a>obs://</strong><em id="i1071973318532"><a name="i1071973318532"></a><a name="i1071973318532"></a>bucketname</em> when downloading all objects in a bucket. Set <strong id="b121504409546"><a name="b121504409546"></a><a name="b121504409546"></a>cloud_url</strong> in the format of <strong id="b0725857165419"><a name="b0725857165419"></a><a name="b0725857165419"></a>obs://</strong><em id="i103736175510"><a name="i103736175510"></a><a name="i103736175510"></a>bucketname</em><strong id="b1049918355516"><a name="b1049918355516"></a><a name="b1049918355516"></a>/</strong><em id="i638387195516"><a name="i638387195516"></a><a name="i638387195516"></a>key</em> when downloading a specified object in a bucket.</p>
</div></div>
</li></ul>
</td>
</tr>
<tr id="row10934256225"><td class="cellrowborder" valign="top" width="18.09180918091809%" headers="mcps1.2.4.1.1 "><p id="p149341562024"><a name="p149341562024"></a><a name="p149341562024"></a>4</p>
</td>
<td class="cellrowborder" valign="top" width="26.82268226822682%" headers="mcps1.2.4.1.2 "><p id="p69344567215"><a name="p69344567215"></a><a name="p69344567215"></a>Bucket status error</p>
</td>
<td class="cellrowborder" valign="top" width="55.08550855085508%" headers="mcps1.2.4.1.3 "><p id="p1693495612219"><a name="p1693495612219"></a><a name="p1693495612219"></a>The specified destination bucket does not exist for uploading a folder by running the <strong id="b155231141161512"><a name="b155231141161512"></a><a name="b155231141161512"></a>cp</strong> command.</p>
</td>
</tr>
<tr id="row1467515181148"><td class="cellrowborder" valign="top" width="18.09180918091809%" headers="mcps1.2.4.1.1 "><p id="p15675718548"><a name="p15675718548"></a><a name="p15675718548"></a>5</p>
</td>
<td class="cellrowborder" valign="top" width="26.82268226822682%" headers="mcps1.2.4.1.2 "><p id="p17675818349"><a name="p17675818349"></a><a name="p17675818349"></a>Initialization error during command execution</p>
</td>
<td class="cellrowborder" valign="top" width="55.08550855085508%" headers="mcps1.2.4.1.3 "><a name="ul145291824329"></a><a name="ul145291824329"></a><ul id="ul145291824329"><li>An error occurs when loading the configuration file.</li><li>Parameter <strong id="b10363845165813"><a name="b10363845165813"></a><a name="b10363845165813"></a>-o</strong> is configured when running the <strong id="b197467480589"><a name="b197467480589"></a><a name="b197467480589"></a>cp</strong> command to upload a folder, but the folder specified by <strong id="b03031601907"><a name="b03031601907"></a><a name="b03031601907"></a>-o</strong> for saving the result lists fails to be created.</li></ul>
</td>
</tr>
<tr id="row911713115716"><td class="cellrowborder" valign="top" width="18.09180918091809%" headers="mcps1.2.4.1.1 "><p id="p811741120713"><a name="p811741120713"></a><a name="p811741120713"></a>6</p>
</td>
<td class="cellrowborder" valign="top" width="26.82268226822682%" headers="mcps1.2.4.1.2 "><p id="p91177111371"><a name="p91177111371"></a><a name="p91177111371"></a>Execution error.</p>
</td>
<td class="cellrowborder" valign="top" width="55.08550855085508%" headers="mcps1.2.4.1.3 "><p id="p141171711975"><a name="p141171711975"></a><a name="p141171711975"></a>When you run the <strong id="b1591032102518"><a name="b1591032102518"></a><a name="b1591032102518"></a>ls</strong> command to query the bucket list, the query fails because the network times out.</p>
</td>
</tr>
<tr id="row1269312411894"><td class="cellrowborder" valign="top" width="18.09180918091809%" headers="mcps1.2.4.1.1 "><p id="p1269354119919"><a name="p1269354119919"></a><a name="p1269354119919"></a>7</p>
</td>
<td class="cellrowborder" valign="top" width="26.82268226822682%" headers="mcps1.2.4.1.2 "><p id="p96937416915"><a name="p96937416915"></a><a name="p96937416915"></a>The operation is not supported.</p>
</td>
<td class="cellrowborder" valign="top" width="55.08550855085508%" headers="mcps1.2.4.1.3 "><p id="p1369318417915"><a name="p1369318417915"></a><a name="p1369318417915"></a>The version of the bucket is not 3.0 when you attempt to change the properties of an object in it by running the <strong id="b724204311317"><a name="b724204311317"></a><a name="b724204311317"></a>chattri</strong> command.</p>
</td>
</tr>
<tr id="row449544513101"><td class="cellrowborder" valign="top" width="18.09180918091809%" headers="mcps1.2.4.1.1 "><p id="p749564519103"><a name="p749564519103"></a><a name="p749564519103"></a>8</p>
</td>
<td class="cellrowborder" valign="top" width="26.82268226822682%" headers="mcps1.2.4.1.2 "><p id="p144951545171013"><a name="p144951545171013"></a><a name="p144951545171013"></a>A batch task succeeded partially.</p>
</td>
<td class="cellrowborder" valign="top" width="55.08550855085508%" headers="mcps1.2.4.1.3 "><p id="p1549514515101"><a name="p1549514515101"></a><a name="p1549514515101"></a>Some objects fail to be downloaded during a batch download by running the <strong id="b14368013182419"><a name="b14368013182419"></a><a name="b14368013182419"></a>cp</strong> command.</p>
</td>
</tr>
<tr id="row1504936141118"><td class="cellrowborder" valign="top" width="18.09180918091809%" headers="mcps1.2.4.1.1 "><p id="p2050413621110"><a name="p2050413621110"></a><a name="p2050413621110"></a>9</p>
</td>
<td class="cellrowborder" valign="top" width="26.82268226822682%" headers="mcps1.2.4.1.2 "><p id="p8504143681116"><a name="p8504143681116"></a><a name="p8504143681116"></a>Interruption error</p>
</td>
<td class="cellrowborder" valign="top" width="55.08550855085508%" headers="mcps1.2.4.1.3 "><p id="p125041362119"><a name="p125041362119"></a><a name="p125041362119"></a>Users press <strong id="b1442612715262"><a name="b1442612715262"></a><a name="b1442612715262"></a>Ctrl</strong> + <strong id="b678713109267"><a name="b678713109267"></a><a name="b678713109267"></a>C</strong> to interrupt the command execution.</p>
</td>
</tr>
<tr id="row59351423116"><td class="cellrowborder" valign="top" width="18.09180918091809%" headers="mcps1.2.4.1.1 "><p id="p14936174241117"><a name="p14936174241117"></a><a name="p14936174241117"></a>-1</p>
</td>
<td class="cellrowborder" valign="top" width="26.82268226822682%" headers="mcps1.2.4.1.2 "><p id="p9936942161114"><a name="p9936942161114"></a><a name="p9936942161114"></a>Unknown error</p>
</td>
<td class="cellrowborder" valign="top" width="55.08550855085508%" headers="mcps1.2.4.1.3 "><p id="p4936842161114"><a name="p4936842161114"></a><a name="p4936842161114"></a>-</p>
</td>
</tr>
</tbody>
</table>

