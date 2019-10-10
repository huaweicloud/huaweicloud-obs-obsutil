# Synchronously Downloading Incremental Objects<a name="EN-US_TOPIC_0150708636"></a>

## Function<a name="section1479112110815"></a>

This function synchronizes all content in the specified path of the source bucket to the target bucket on OBS, ensuring that the content is consistent between the specified path of the source bucket and the target bucket. Incremental synchronization has the following meanings: 1\) Increment: Compare the source object with the target file and download only the source object that has changes. 2\) Synchronization: After the command is executed, ensure that the specified path of the source bucket is a subset of the local target path. That is, any object in the specified path of the source bucket has its corresponding file in the local target path.

>![](public_sys-resources/icon-notice.gif) **NOTICE:**   
>-   Do not change the source objects in the OBS bucket during synchronization. Otherwise, the synchronization may fail or data may be inconsistent.  
>-   If the storage class of the object to be copied is  **cold**, you must restore the object to be downloaded first. Otherwise, the download fails.  
>-   Each object can be synchronously downloaded only when it does not exist in the local path, its size is different from the namesake one in the local path, or it has the latest modification time.  

## Command Line Structure<a name="section19587175622519"></a>

-   In Windows

    ```
    obsutil sync obs://bucket[/key] folder_url [-tempFileDir=xxx] [-dryRun] [-vlength] [-vmd5] [-j=1] [-p=1] [-threshold=52428800] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]
    ```

-   In Linux or macOS

    ```
    ./obsutil sync obs://bucket[/key] folder_url [-tempFileDir=xxx] [-dryRun] [-vlength] [-vmd5] [-j=1] [-p=1] [-threshold=52428800] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]
    ```


## Parameter Description<a name="section175878561255"></a>

<a name="table1155454517538"></a>
<table><thead align="left"><tr id="row6558154516531"><th class="cellrowborder" valign="top" width="18%" id="mcps1.1.4.1.1"><p id="p1656117451534"><a name="p1656117451534"></a><a name="p1656117451534"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="27%" id="mcps1.1.4.1.2"><p id="p155651545195310"><a name="p155651545195310"></a><a name="p155651545195310"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="55.00000000000001%" id="mcps1.1.4.1.3"><p id="p165678456530"><a name="p165678456530"></a><a name="p165678456530"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row195751145115315"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p1576945115318"><a name="p1576945115318"></a><a name="p1576945115318"></a>folder_url</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p75772453537"><a name="p75772453537"></a><a name="p75772453537"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p12579194519531"><a name="p12579194519531"></a><a name="p12579194519531"></a>Local folder path</p>
</td>
</tr>
<tr id="row1157912451534"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p058211459536"><a name="p058211459536"></a><a name="p058211459536"></a>bucket</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p5582164575310"><a name="p5582164575310"></a><a name="p5582164575310"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p158414515316"><a name="p158414515316"></a><a name="p158414515316"></a>Bucket name</p>
</td>
</tr>
<tr id="row1584114513535"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p17585144585310"><a name="p17585144585310"></a><a name="p17585144585310"></a>key</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p10783112824619"><a name="p10783112824619"></a><a name="p10783112824619"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p0588134575318"><a name="p0588134575318"></a><a name="p0588134575318"></a>Indicates the name prefix of objects to be synchronously downloaded.</p>
<p id="p115612190125"><a name="p115612190125"></a><a name="p115612190125"></a>The rules are as follows:</p>
<a name="ul15158191991211"></a><a name="ul15158191991211"></a><ul id="ul15158191991211"><li>If this parameter is left blank, all files in the folder specified by <strong id="b208353894311"><a name="b208353894311"></a><a name="b208353894311"></a>folder_url</strong> are the same as all objects in the bucket.</li><li>If this parameter is configured, all files in the folder specified by <strong id="b1872818114811"><a name="b1872818114811"></a><a name="b1872818114811"></a>folder_url</strong> are the same as the objects whose name prefix is the configured value in the bucket.</li></ul>
<div class="note" id="note14528254193716"><a name="note14528254193716"></a><a name="note14528254193716"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul54182557214"></a><a name="ul54182557214"></a><ul id="ul54182557214"><li>If the value of this parameter does not end with a slash (/), the obsutil tool automatically adds a slash (/) at the end of the configured value as the object name prefix.</li><li>For details about how to use this parameter, see <a href="synchronous-download-examples.md">Synchronous Download Examples</a>.</li></ul>
</div></div>
</td>
</tr>
<tr id="row187592010131217"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p276284012715"><a name="p276284012715"></a><a name="p276284012715"></a>tempFileDir</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p122055912103"><a name="p122055912103"></a><a name="p122055912103"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1476264014710"><a name="p1476264014710"></a><a name="p1476264014710"></a>Indicates the directory for storing temporary files during synchronous download. The default value is the value of <strong id="b498311817912"><a name="b498311817912"></a><a name="b498311817912"></a>defaultTempFileDir</strong> in the configuration file.</p>
<div class="note" id="note1564311143011"><a name="note1564311143011"></a><a name="note1564311143011"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul15994748142619"></a><a name="ul15994748142619"></a><ul id="ul15994748142619"><li>Temporary files generated during multipart download are stored in this directory. Therefore, ensure that the user who executes obsutil has the write permission on the path.</li><li>The available space of the partition where the path is located must be greater than the size of the objects to be downloaded.</li></ul>
</div></div>
</td>
</tr>
<tr id="row35998457538"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p1537517155414"><a name="p1537517155414"></a><a name="p1537517155414"></a>dryRun</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p13376216548"><a name="p13376216548"></a><a name="p13376216548"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1137981175411"><a name="p1137981175411"></a><a name="p1137981175411"></a>Conducts a dry run.</p>
</td>
</tr>
<tr id="row9603194595311"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p1605745105311"><a name="p1605745105311"></a><a name="p1605745105311"></a>vlength</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p12605114512534"><a name="p12605114512534"></a><a name="p12605114512534"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1660724595316"><a name="p1660724595316"></a><a name="p1660724595316"></a>Checks whether the sizes of the local files are the same as those of the objects in the bucket after the download is complete.</p>
</td>
</tr>
<tr id="row360812456535"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p136091945105315"><a name="p136091945105315"></a><a name="p136091945105315"></a>vmd5</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p461011452539"><a name="p461011452539"></a><a name="p461011452539"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p19611745185310"><a name="p19611745185310"></a><a name="p19611745185310"></a>Checks whether MD5 values of the local files are the same as those of the objects in the bucket after the download is complete.</p>
<div class="note" id="note185321026122811"><a name="note185321026122811"></a><a name="note185321026122811"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p12303311112719"><a name="p12303311112719"></a><a name="p12303311112719"></a>Objects in the bucket must contain metadata <strong id="b1056814153246"><a name="b1056814153246"></a><a name="b1056814153246"></a>x-obs-md5chksum</strong>. Otherwise, MD5 verification will be skipped.</p>
</div></div>
</td>
</tr>
<tr id="row36131445115310"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p14613194514535"><a name="p14613194514535"></a><a name="p14613194514535"></a>p</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p15615164515317"><a name="p15615164515317"></a><a name="p15615164515317"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1161744513534"><a name="p1161744513534"></a><a name="p1161744513534"></a>Indicates the maximum number of concurrent multipart download tasks when downloading an object. The default value is the value of <strong id="b197311503403"><a name="b197311503403"></a><a name="b197311503403"></a>defaultParallels</strong> in the configuration file.</p>
</td>
</tr>
<tr id="row14617144520537"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p17619545185314"><a name="p17619545185314"></a><a name="p17619545185314"></a>threshold</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p9620645185311"><a name="p9620645185311"></a><a name="p9620645185311"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p262174575315"><a name="p262174575315"></a><a name="p262174575315"></a>Indicates the threshold for enabling multipart download, in bytes. The default value is the value of <strong id="b19309194493311"><a name="b19309194493311"></a><a name="b19309194493311"></a>defaultBigfileThreshold</strong> in the configuration file.</p>
<div class="note" id="note94444118613"><a name="note94444118613"></a><a name="note94444118613"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul182059695115"></a><a name="ul182059695115"></a><ul id="ul182059695115"><li>If the size of the object to be downloaded is smaller than the threshold, download the object directly. If not, a multipart download is required.</li><li>If you download an object directly, no part record is generated, and resumable transmission is not supported.</li><li>This value can contain a capacity unit. For example, <strong id="b18188163013710"><a name="b18188163013710"></a><a name="b18188163013710"></a>1 MB</strong> indicates 1048576 bytes.</li></ul>
</div></div>
</td>
</tr>
<tr id="row1363634595315"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p3637164545318"><a name="p3637164545318"></a><a name="p3637164545318"></a>ps</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p1264164517532"><a name="p1264164517532"></a><a name="p1264164517532"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p10641184516531"><a name="p10641184516531"></a><a name="p10641184516531"></a>Indicates the size of each part in a multipart download task, in bytes. The default value is the value of <strong id="b71116458717"><a name="b71116458717"></a><a name="b71116458717"></a>defaultPartSize</strong> in the configuration file.</p>
<div class="note" id="note1175045524310"><a name="note1175045524310"></a><a name="note1175045524310"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul334662519451"></a><a name="ul334662519451"></a><ul id="ul334662519451"><li>This value can contain a capacity unit. For example, <strong id="b4127164695113"><a name="b4127164695113"></a><a name="b4127164695113"></a>1 MB</strong> indicates 1048576 bytes.</li><li>The parameter can be set to <strong id="b16689155095119"><a name="b16689155095119"></a><a name="b16689155095119"></a>auto</strong>. In this case, obsutil automatically sets the part size for each multipart task based on the source object size.</li></ul>
</div></div>
</td>
</tr>
<tr id="row196461145155314"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p1764874511530"><a name="p1764874511530"></a><a name="p1764874511530"></a>cpd</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p206482455537"><a name="p206482455537"></a><a name="p206482455537"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1965018458536"><a name="p1965018458536"></a><a name="p1965018458536"></a>Indicates the folder where the part records reside. The default value is <strong id="b191878914259"><a name="b191878914259"></a><a name="b191878914259"></a>.obsutil_checkpoint</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note519835512711"><a name="note519835512711"></a><a name="note519835512711"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p118866577266"><a name="p118866577266"></a><a name="p118866577266"></a>A part record is generated during a multipart download and saved to the <strong id="b356213551770"><a name="b356213551770"></a><a name="b356213551770"></a>down</strong> subfolder. After the download succeeds, its part record is deleted automatically. If the download fails or is suspended, the system attempts to resume the task according to its part record when you perform the download the next time.</p>
</div></div>
</td>
</tr>
<tr id="row1566014510533"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p3662204585313"><a name="p3662204585313"></a><a name="p3662204585313"></a>j</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p19305175612615"><a name="p19305175612615"></a><a name="p19305175612615"></a>Optional for downloading objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p17665114555311"><a name="p17665114555311"></a><a name="p17665114555311"></a>Indicates the maximum number of concurrent tasks for downloading objects synchronously. The default value is the value of <strong id="b115503229418"><a name="b115503229418"></a><a name="b115503229418"></a>defaultJobs</strong> in the configuration file.</p>
<div class="note" id="note891964620819"><a name="note891964620819"></a><a name="note891964620819"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p1091964618820"><a name="p1091964618820"></a><a name="p1091964618820"></a>The value is ensured to be greater than or equal to 1.</p>
</div></div>
</td>
</tr>
<tr id="row109762261097"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p156173113915"><a name="p156173113915"></a><a name="p156173113915"></a>exclude</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p1930885652611"><a name="p1930885652611"></a><a name="p1930885652611"></a>Optional for downloading objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p16666614131519"><a name="p16666614131519"></a><a name="p16666614131519"></a>Indicates the matching patterns of source objects that are excluded, for example: <strong id="b03706457276"><a name="b03706457276"></a><a name="b03706457276"></a>*.txt</strong>.</p>
<div class="note" id="note145284716208"><a name="note145284716208"></a><a name="note145284716208"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul05604713204"></a><a name="ul05604713204"></a><ul id="ul05604713204"><li>The asterisk (*) represents any group of characters, and the question mark (?) represents any single character. For instance, <strong id="b19823146112712"><a name="b19823146112712"></a><a name="b19823146112712"></a>abc*.txt</strong> indicates any file whose name starts with <strong id="b18825846182714"><a name="b18825846182714"></a><a name="b18825846182714"></a>abc</strong> and ends with <strong id="b882720465272"><a name="b882720465272"></a><a name="b882720465272"></a>.txt</strong>.</li><li>You can use <strong id="b12211248182714"><a name="b12211248182714"></a><a name="b12211248182714"></a>\*</strong> to represent <strong id="b1322234832717"><a name="b1322234832717"></a><a name="b1322234832717"></a>*</strong> and <strong id="b1122311486272"><a name="b1122311486272"></a><a name="b1122311486272"></a>\?</strong> to represent <strong id="b122416488276"><a name="b122416488276"></a><a name="b122416488276"></a>?</strong>.</li><li>If the name of the object to be downloaded matches the value of this parameter, the object is skipped.</li></ul>
</div></div>
<div class="notice" id="note179117549207"><a name="note179117549207"></a><a name="note179117549207"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul877892116516"></a><a name="ul877892116516"></a><ul id="ul877892116516"><li>You are advised to use quotation marks for the matching pattern to prevent special characters from being escaped by the OS and leading to unexpected results. Use single quotation marks for Linux or macOS and quotation marks for Windows.</li><li>The matching pattern applies to the absolute path of an object, including the object name prefix and object name starting from the root directory. For example, if the path of an object in the bucket is <strong id="b1479674314319"><a name="b1479674314319"></a><a name="b1479674314319"></a>obs://bucket/src1/src2/test.txt</strong>, then the absolute path of the object is <strong id="b9796443123114"><a name="b9796443123114"></a><a name="b9796443123114"></a>src1/src2/test.txt</strong>.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li></ul>
</div></div>
</td>
</tr>
<tr id="row20666154513530"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p186661845145315"><a name="p186661845145315"></a><a name="p186661845145315"></a>include</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p966784516533"><a name="p966784516533"></a><a name="p966784516533"></a>Optional for downloading objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p37071730153119"><a name="p37071730153119"></a><a name="p37071730153119"></a>Indicates the matching patterns of source objects that are included, for example: <strong id="b153581454102715"><a name="b153581454102715"></a><a name="b153581454102715"></a>*.jpg</strong>.</p>
<div class="note" id="note195168716220"><a name="note195168716220"></a><a name="note195168716220"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul752013715229"></a><a name="ul752013715229"></a><ul id="ul752013715229"><li>The asterisk (*) represents any group of characters, and the question mark (?) represents any single character.</li><li>You can use <strong id="b11288574273"><a name="b11288574273"></a><a name="b11288574273"></a>\*</strong> to represent <strong id="b7129257182719"><a name="b7129257182719"></a><a name="b7129257182719"></a>*</strong> and <strong id="b11130757112717"><a name="b11130757112717"></a><a name="b11130757112717"></a>\?</strong> to represent <strong id="b11314574274"><a name="b11314574274"></a><a name="b11314574274"></a>?</strong>.</li><li>Only after identifying that the name of the file to be downloaded does not match the value of <strong id="b517175915272"><a name="b517175915272"></a><a name="b517175915272"></a>exclude</strong>, the system checks whether the file name matches the value of this parameter. If yes, the file is downloaded. If not, the file is skipped.</li></ul>
</div></div>
<div class="notice" id="note9270217202212"><a name="note9270217202212"></a><a name="note9270217202212"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul14807204120586"></a><a name="ul14807204120586"></a><ul id="ul14807204120586"><li>You are advised to use quotation marks for the matching pattern to prevent special characters from being escaped by the OS and leading to unexpected results. Use single quotation marks for Linux or macOS and quotation marks for Windows.</li><li>The matching pattern applies to the absolute path of an object, including the object name prefix and object name starting from the root directory. For example, if the path of an object in the bucket is <strong id="b6347174663115"><a name="b6347174663115"></a><a name="b6347174663115"></a>obs://bucket/src1/src2/test.txt</strong>, then the absolute path of the object is <strong id="b1834884613314"><a name="b1834884613314"></a><a name="b1834884613314"></a>src1/src2/test.txt</strong>.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li></ul>
</div></div>
</td>
</tr>
<tr id="row936321020457"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p496314014481"><a name="p496314014481"></a><a name="p496314014481"></a>timeRange</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p1664510223494"><a name="p1664510223494"></a><a name="p1664510223494"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p16963940134813"><a name="p16963940134813"></a><a name="p16963940134813"></a>Indicates the time range matching pattern when synchronously downloading objects. Only objects whose last modification time is within the configured time range are downloaded.</p>
<p id="p1754153575414"><a name="p1754153575414"></a><a name="p1754153575414"></a>This pattern has a lower priority than the object matching patterns (<strong id="b445318471394"><a name="b445318471394"></a><a name="b445318471394"></a>exclude</strong>/<strong id="b1454174713911"><a name="b1454174713911"></a><a name="b1454174713911"></a>include</strong>). That is, the time range matching pattern is executed after the configured object matching patterns.</p>
<div class="note" id="note8766915165919"><a name="note8766915165919"></a><a name="note8766915165919"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul19272430145915"></a><a name="ul19272430145915"></a><ul id="ul19272430145915"><li>The matching time range is represented in <em id="i15681115018316"><a name="i15681115018316"></a><a name="i15681115018316"></a>time1</em><strong id="b9682450123118"><a name="b9682450123118"></a><a name="b9682450123118"></a>-</strong><em id="i1868365011310"><a name="i1868365011310"></a><a name="i1868365011310"></a>time2</em>, where <em id="i868365012316"><a name="i868365012316"></a><a name="i868365012316"></a>time1</em> must be earlier than or the same as <em id="i1683125018315"><a name="i1683125018315"></a><a name="i1683125018315"></a>time2</em>. The time format is <em id="i11684650103114"><a name="i11684650103114"></a><a name="i11684650103114"></a>yyyyMMddHHmmss</em>.</li><li>Automatic formatting is supported. For example, yyyyMMdd is equivalent to yyyyMMdd000000, and yyyyMM is equivalent to yyyyMM01000000.</li><li>If this parameter is set to <strong id="b18160175318312"><a name="b18160175318312"></a><a name="b18160175318312"></a>*-</strong><em id="i1516195343118"><a name="i1516195343118"></a><a name="i1516195343118"></a>time2</em>, all files whose last modification time is earlier than <em id="i1162155315316"><a name="i1162155315316"></a><a name="i1162155315316"></a>time2</em> are matched. If it is set to <em id="i61625530319"><a name="i61625530319"></a><a name="i61625530319"></a>time1</em><strong id="b191638533315"><a name="b191638533315"></a><a name="b191638533315"></a>-*</strong>, all files whose last modification time is later than <em id="i19163553203113"><a name="i19163553203113"></a><a name="i19163553203113"></a>time1</em> are matched.</li></ul>
</div></div>
<div class="notice" id="note1078981817591"><a name="note1078981817591"></a><a name="note1078981817591"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul881073612597"></a><a name="ul881073612597"></a><ul id="ul881073612597"><li>Time in the matching pattern is the UTC time.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li></ul>
</div></div>
</td>
</tr>
<tr id="row16694195516812"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p10122175516"><a name="p10122175516"></a><a name="p10122175516"></a>mf</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p1562018121550"><a name="p1562018121550"></a><a name="p1562018121550"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p31226718512"><a name="p31226718512"></a><a name="p31226718512"></a>Indicates that the name matching pattern (<strong id="b819818124514"><a name="b819818124514"></a><a name="b819818124514"></a>include</strong> or <strong id="b1519918124511"><a name="b1519918124511"></a><a name="b1519918124511"></a>exclude</strong>) and the time matching pattern (<strong id="b121993125515"><a name="b121993125515"></a><a name="b121993125515"></a>timeRange</strong>) also take effect on objects whose names end with a slash (/).</p>
<p id="p15797105118514"><a name="p15797105118514"></a><a name="p15797105118514"></a></p>
</td>
</tr>
<tr id="row109093549288"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p143226572287"><a name="p143226572287"></a><a name="p143226572287"></a>o</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p232405732811"><a name="p232405732811"></a><a name="p232405732811"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1632615718282"><a name="p1632615718282"></a><a name="p1632615718282"></a>Indicates the folder where operation result lists reside. After the command is executed, result lists (possibly including success, failure, and warning files) are generated in the folder. The default value is <strong id="b14198182512519"><a name="b14198182512519"></a><a name="b14198182512519"></a>.obsutil_output</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note1416817409247"><a name="note1416817409247"></a><a name="note1416817409247"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul101190347408"></a><a name="ul101190347408"></a><ul id="ul101190347408"><li>The naming rule for result lists is as follows: <strong id="b12140172782415"><a name="b12140172782415"></a><a name="b12140172782415"></a>sync_{succeed  | failed | warning}_report_</strong><em id="i151415274241"><a name="i151415274241"></a><a name="i151415274241"></a>time</em><strong id="b1614214271245"><a name="b1614214271245"></a><a name="b1614214271245"></a>_TaskId.txt</strong></li><li>By default, the maximum size of a single result list is 30 MB and the maximum number of result lists that can be retained is 1024. You can set the maximum size and number by configuring <strong id="b17707141163819"><a name="b17707141163819"></a><a name="b17707141163819"></a>recordMaxLogSize</strong> and <strong id="b14708174143820"><a name="b14708174143820"></a><a name="b14708174143820"></a>recordBackups</strong> in the configuration file.</li></ul>
</div></div>
</td>
</tr>
<tr id="row878983516371"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Response<a name="section6926520122416"></a>

Refer to  [Response](uploading-an-object.md#section6926520122416)  for uploading an object.

## Running Example<a name="section1059235632510"></a>

-   Take the Windows OS as an example. Run the  **obsutil sync obs://bucket-test/temp d:\\ temp**  command to download objects synchronously.

```
obsutil sync obs://bucket-test/temp d:\temp

Parallel:      3                   Jobs:          3
Threshold:     524288000           PartSize:      5242880
Exclude:                           Include:
VerifyLength:  false               VerifyMd5:     false
CheckpointDir: xxxx
OutputDir: xxxx

[======================================================] 100.00% 155.59 KB/s 0s
Succeed count is:   6         Failed count is:    0
Metrics [max cost:153 ms, min cost:129 ms, average cost:92.00 ms, average tps:17.86]
Task id is: 3066a4b0-4d21-4929-bb84-4829c32cbd0f
```

-   For more examples, see  [Synchronous Download Examples](synchronous-download-examples.md).

