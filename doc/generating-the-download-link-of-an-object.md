# Generating the Download Link of an Object<a name="EN-US_TOPIC_0160351846"></a>

## Function<a name="section1479112110815"></a>

You can use this command to generate the download link of a specified object in a bucket or generate the download links of objects in a bucket in batches by object name prefix.

## Command Line Structure<a name="section1220752192216"></a>

-   In Windows
    -   Generating the download link of a single object

        ```
        obsutil sign obs://bucket/key [-e=300] [-config=xxx]
        ```

    -   Generating the download links of objects in batches by object name prefix

        ```
        obsutil sign obs://bucket[/key] -r [-e=300] [-timeRange=time1-time2] [-include=*.xxx] [-exclude=*.xxx] [-o=xxx] [-config=xxx]
        ```


-   In Linux or macOS
    -   Generating the download link of a single object

        ```
        ./obsutil sign obs://bucket/key [-e=300] [-config=xxx]
        ```

    -   Generating the download links of objects in batches by object name prefix

        ```
        ./obsutil sign obs://bucket[/key] -r [-e=300] [-timeRange=time1-time2] [-include=*.xxx] [-exclude=*.xxx] [-o=xxx] [-config=xxx]
        ```



## Parameter Description<a name="section6559191102418"></a>

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="13%" id="mcps1.1.4.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="18%" id="mcps1.1.4.1.2"><p id="p15137125919108"><a name="p15137125919108"></a><a name="p15137125919108"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="69%" id="mcps1.1.4.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row1096715662419"><td class="cellrowborder" valign="top" width="13%" headers="mcps1.1.4.1.1 "><p id="p11357141182414"><a name="p11357141182414"></a><a name="p11357141182414"></a>bucket</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.2 "><p id="p12357511172416"><a name="p12357511172416"></a><a name="p12357511172416"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="69%" headers="mcps1.1.4.1.3 "><p id="p1435717116242"><a name="p1435717116242"></a><a name="p1435717116242"></a>Bucket name</p>
</td>
</tr>
<tr id="row129424415242"><td class="cellrowborder" valign="top" width="13%" headers="mcps1.1.4.1.1 "><p id="p11944643141811"><a name="p11944643141811"></a><a name="p11944643141811"></a>key</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.2 "><p id="p19639114118186"><a name="p19639114118186"></a><a name="p19639114118186"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="69%" headers="mcps1.1.4.1.3 "><p id="p16965474526"><a name="p16965474526"></a><a name="p16965474526"></a>Object name used for generating the download link of a single object, or object name prefix used for generating download links of objects in batches</p>
</td>
</tr>
<tr id="row1320919474175"><td class="cellrowborder" valign="top" width="13%" headers="mcps1.1.4.1.1 "><p id="p1720914714179"><a name="p1720914714179"></a><a name="p1720914714179"></a>e</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.2 "><p id="p16209347141715"><a name="p16209347141715"></a><a name="p16209347141715"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="69%" headers="mcps1.1.4.1.3 "><p id="p17209347151710"><a name="p17209347151710"></a><a name="p17209347151710"></a>Validity period of the generated download links of objects, in seconds. Minimum value: 60s. Default value: 300s</p>
</td>
</tr>
<tr id="row1530918494177"><td class="cellrowborder" valign="top" width="13%" headers="mcps1.1.4.1.1 "><p id="p15309164911712"><a name="p15309164911712"></a><a name="p15309164911712"></a>r</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.2 "><p id="p7309749181710"><a name="p7309749181710"></a><a name="p7309749181710"></a>Mandatory when generating download links of objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="69%" headers="mcps1.1.4.1.3 "><p id="p15309114931714"><a name="p15309114931714"></a><a name="p15309114931714"></a>Generates the download links of objects in batches by a specified object name prefix.</p>
</td>
</tr>
<tr id="row982613333425"><td class="cellrowborder" valign="top" width="13%" headers="mcps1.1.4.1.1 "><p id="p156173113915"><a name="p156173113915"></a><a name="p156173113915"></a>exclude</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.2 "><p id="p1930885652611"><a name="p1930885652611"></a><a name="p1930885652611"></a>Optional when generating download links of objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="69%" headers="mcps1.1.4.1.3 "><p id="p16666614131519"><a name="p16666614131519"></a><a name="p16666614131519"></a>Indicates the matching patterns of objects that are excluded, for example: <strong id="b141934383516"><a name="b141934383516"></a><a name="b141934383516"></a>*.txt</strong>.</p>
<div class="note" id="note145284716208"><a name="note145284716208"></a><a name="note145284716208"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul05604713204"></a><a name="ul05604713204"></a><ul id="ul05604713204"><li>The asterisk (*) represents any group of characters, and the question mark (?) represents any single character. For instance, <strong id="b960954194116"><a name="b960954194116"></a><a name="b960954194116"></a>abc*.txt</strong> indicates any file whose name starts with <strong id="b196913548413"><a name="b196913548413"></a><a name="b196913548413"></a>abc</strong> and ends with <strong id="b1070185411417"><a name="b1070185411417"></a><a name="b1070185411417"></a>.txt</strong>.</li><li>You can use <strong id="b34091756134117"><a name="b34091756134117"></a><a name="b34091756134117"></a>\*</strong> to represent <strong id="b19410165616411"><a name="b19410165616411"></a><a name="b19410165616411"></a>*</strong> and <strong id="b154111456204116"><a name="b154111456204116"></a><a name="b154111456204116"></a>\?</strong> to represent <strong id="b241265644119"><a name="b241265644119"></a><a name="b241265644119"></a>?</strong>.</li><li>If the name of the object to be downloaded matches the value of this parameter, the object is skipped.</li></ul>
</div></div>
<div class="notice" id="note179117549207"><a name="note179117549207"></a><a name="note179117549207"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul877892116516"></a><a name="ul877892116516"></a><ul id="ul877892116516"><li>You are advised to use quotation marks for the matching pattern to prevent special characters from being escaped by the OS and leading to unexpected results. Use single quotation marks for Linux or macOS and quotation marks for Windows.</li><li>The matching pattern applies to the absolute path of an object, including the object name prefix and object name starting from the root directory. For example, if the path of an object in the bucket is <strong id="b639119353111"><a name="b639119353111"></a><a name="b639119353111"></a>obs://bucket/src1/src2/test.txt</strong>, then the absolute path of the object is <strong id="b123921738315"><a name="b123921738315"></a><a name="b123921738315"></a>src1/src2/test.txt</strong>.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li></ul>
</div></div>
</td>
</tr>
<tr id="row17153224213"><td class="cellrowborder" valign="top" width="13%" headers="mcps1.1.4.1.1 "><p id="p186661845145315"><a name="p186661845145315"></a><a name="p186661845145315"></a>include</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.2 "><p id="p966784516533"><a name="p966784516533"></a><a name="p966784516533"></a>Optional when generating download links of objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="69%" headers="mcps1.1.4.1.3 "><p id="p37071730153119"><a name="p37071730153119"></a><a name="p37071730153119"></a>Indicates the matching patterns of objects that are included, for example: <strong id="b16049125112"><a name="b16049125112"></a><a name="b16049125112"></a>*.jpg</strong>.</p>
<div class="note" id="note195168716220"><a name="note195168716220"></a><a name="note195168716220"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul752013715229"></a><a name="ul752013715229"></a><ul id="ul752013715229"><li>The asterisk (*) represents any group of characters, and the question mark (?) represents any single character.</li><li>You can use <strong id="b987530144318"><a name="b987530144318"></a><a name="b987530144318"></a>\*</strong> to represent <strong id="b18877180184311"><a name="b18877180184311"></a><a name="b18877180184311"></a>*</strong> and <strong id="b1287813084316"><a name="b1287813084316"></a><a name="b1287813084316"></a>\?</strong> to represent <strong id="b287910134320"><a name="b287910134320"></a><a name="b287910134320"></a>?</strong>.</li><li>Only after identifying that the name of the file to be downloaded does not match the value of <strong id="b613133154313"><a name="b613133154313"></a><a name="b613133154313"></a>exclude</strong>, the system checks whether the file name matches the value of this parameter. If yes, the file is downloaded. If not, the file is skipped.</li></ul>
</div></div>
<div class="notice" id="note9270217202212"><a name="note9270217202212"></a><a name="note9270217202212"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul1329413124584"></a><a name="ul1329413124584"></a><ul id="ul1329413124584"><li>You are advised to use quotation marks for the matching pattern to prevent special characters from being escaped by the OS and leading to unexpected results. Use single quotation marks for Linux or macOS and quotation marks for Windows.</li><li>The matching pattern applies to the absolute path of an object, including the object name prefix and object name starting from the root directory. For example, if the path of an object in the bucket is <strong id="b81280613316"><a name="b81280613316"></a><a name="b81280613316"></a>obs://bucket/src1/src2/test.txt</strong>, then the absolute path of the object is <strong id="b6129365319"><a name="b6129365319"></a><a name="b6129365319"></a>src1/src2/test.txt</strong>.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li></ul>
</div></div>
</td>
</tr>
<tr id="row143423575435"><td class="cellrowborder" valign="top" width="13%" headers="mcps1.1.4.1.1 "><p id="p496314014481"><a name="p496314014481"></a><a name="p496314014481"></a>timeRange</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.2 "><p id="p1664510223494"><a name="p1664510223494"></a><a name="p1664510223494"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="69%" headers="mcps1.1.4.1.3 "><p id="p16963940134813"><a name="p16963940134813"></a><a name="p16963940134813"></a>Indicates the time range matching pattern when generating download links of objects. Only the download links of objects whose last modification time is within the configured time range are generated.</p>
<p id="p2816559125317"><a name="p2816559125317"></a><a name="p2816559125317"></a>This pattern has a lower priority than the object matching patterns (<strong id="b51820100612"><a name="b51820100612"></a><a name="b51820100612"></a>exclude</strong>/<strong id="b132018108614"><a name="b132018108614"></a><a name="b132018108614"></a>include</strong>). That is, the time range matching pattern is executed after the configured object matching patterns.</p>
<div class="note" id="note8766915165919"><a name="note8766915165919"></a><a name="note8766915165919"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul19272430145915"></a><a name="ul19272430145915"></a><ul id="ul19272430145915"><li>Time in the matching pattern is the UTC time.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li><li>The matching time range is represented in <em id="i12915187173118"><a name="i12915187173118"></a><a name="i12915187173118"></a>time1</em><strong id="b169157716311"><a name="b169157716311"></a><a name="b169157716311"></a>-</strong><em id="i189161733112"><a name="i189161733112"></a><a name="i189161733112"></a>time2</em>, where <em id="i8917774312"><a name="i8917774312"></a><a name="i8917774312"></a>time1</em> must be earlier than or the same as <em id="i11917375311"><a name="i11917375311"></a><a name="i11917375311"></a>time2</em>. The time format is <em id="i5918579310"><a name="i5918579310"></a><a name="i5918579310"></a>yyyyMMddHHmmss</em>.</li><li>Automatic formatting is supported. For example, yyyyMMdd is equivalent to yyyyMMdd000000, and yyyyMM is equivalent to yyyyMM01000000.</li><li>If this parameter is set to <strong id="b179177993112"><a name="b179177993112"></a><a name="b179177993112"></a>*-</strong><em id="i7918398313"><a name="i7918398313"></a><a name="i7918398313"></a>time2</em>, all files whose last modification time is earlier than <em id="i15918159123117"><a name="i15918159123117"></a><a name="i15918159123117"></a>time2</em> are matched. If it is set to <em id="i109195963111"><a name="i109195963111"></a><a name="i109195963111"></a>time1</em><strong id="b1391999143115"><a name="b1391999143115"></a><a name="b1391999143115"></a>-*</strong>, all files whose last modification time is later than <em id="i1792019173116"><a name="i1792019173116"></a><a name="i1792019173116"></a>time1</em> are matched.</li></ul>
</div></div>
</td>
</tr>
<tr id="row0253155181712"><td class="cellrowborder" valign="top" width="13%" headers="mcps1.1.4.1.1 "><p id="p102531251131712"><a name="p102531251131712"></a><a name="p102531251131712"></a>o</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.2 "><p id="p92538516175"><a name="p92538516175"></a><a name="p92538516175"></a>Optional when generating download links of objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="69%" headers="mcps1.1.4.1.3 "><p id="p1632615718282"><a name="p1632615718282"></a><a name="p1632615718282"></a>Indicates the folder where operation result lists reside. After the command is executed, result lists (possibly including success and failure files) are generated in the folder. The default value is <strong id="b228518145437"><a name="b228518145437"></a><a name="b228518145437"></a>.obsutil_output</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note1416817409247"><a name="note1416817409247"></a><a name="note1416817409247"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul101190347408"></a><a name="ul101190347408"></a><ul id="ul101190347408"><li>The naming rule for result lists is as follows: <strong id="b78111648155812"><a name="b78111648155812"></a><a name="b78111648155812"></a>sign_{succeed  | failed}_report_</strong><em id="i381234820584"><a name="i381234820584"></a><a name="i381234820584"></a>time</em><strong id="b1981244865819"><a name="b1981244865819"></a><a name="b1981244865819"></a>_TaskId.txt</strong><p id="li412103664016p0"><a name="li412103664016p0"></a><a name="li412103664016p0"></a>By default, the maximum size of a single result list is 30 MB and the maximum number of result lists that can be retained is 1024. You can set the maximum size and number by configuring <strong id="b194466589359"><a name="b194466589359"></a><a name="b194466589359"></a>recordMaxLogSize</strong> and <strong id="b1447358113513"><a name="b1447358113513"></a><a name="b1447358113513"></a>recordBackups</strong> in the configuration file.</p>
</li></ul>
</div></div>
</td>
</tr>
<tr id="row645016429369"><td class="cellrowborder" valign="top" width="13%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="69%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Running Example<a name="section15899161919244"></a>

-   Take the Windows OS as an example. Run the  **obsutil sign obs://bucket-test/test.txt**  command to generate the download link of a single object.

```
obsutil sign obs://bucket-test/test.txt

Download url of [obs://bucket-test/test.txt] is:
  http://your-endpoint/bucket-test/test.txt?AccessKeyId=xxxx&Expires=1552548758&Signature=xxxx
```

