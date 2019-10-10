# Resuming a Failed Download Task<a name="EN-US_TOPIC_0145591178"></a>

## Function<a name="section1479112110815"></a>

You can use this command to resume a failed download task based on the task ID.

## Command Line Structure<a name="section19587175622519"></a>

-   In Windows

    ```
    obsutil cp -recover=xxx [-dryRun] [-tempFileDir=xxx] [-f] [-u] [-vlength] [-vmd5] [-j=1] [-p=1] [-threshold=52428800] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]
    ```

-   In Linux or macOS

    ```
    ./obsutil cp -recover=xxx [-dryRun] [-tempFileDir=xxx] [-f] [-u] [-vlength] [-vmd5] [-j=1] [-p=1] [-threshold=52428800] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]
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
<tbody><tr id="row43660596231"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p28307351476"><a name="p28307351476"></a><a name="p28307351476"></a>recover</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p198304351672"><a name="p198304351672"></a><a name="p198304351672"></a>Mandatory (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p188308358715"><a name="p188308358715"></a><a name="p188308358715"></a>ID of the download task to be resumed</p>
<div class="note" id="note10332165610814"><a name="note10332165610814"></a><a name="note10332165610814"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul625425894"></a><a name="ul625425894"></a><ul id="ul625425894"><li>You can obtain the task ID after a download task is complete, or query it based on the file name of the operation result list, which is the 36 characters excluding the suffix <strong id="b1948597115915"><a name="b1948597115915"></a><a name="b1948597115915"></a>.txt</strong> in the file name.</li><li>You can locate the download task to be resumed in the directory where the result lists reside. For details about the directory of the result lists, see additional parameter <strong id="b1213073355816"><a name="b1213073355816"></a><a name="b1213073355816"></a>o</strong>.</li></ul>
</div></div>
</td>
</tr>
<tr id="row07629408710"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p276284012715"><a name="p276284012715"></a><a name="p276284012715"></a>tempFileDir</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p122055912103"><a name="p122055912103"></a><a name="p122055912103"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1476264014710"><a name="p1476264014710"></a><a name="p1476264014710"></a>Indicates the directory for storing temporary files during download. The default value is the value of <strong id="b117932163503"><a name="b117932163503"></a><a name="b117932163503"></a>defaultTempFileDir</strong> in the configuration file.</p>
<div class="note" id="note369290203119"><a name="note369290203119"></a><a name="note369290203119"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul15994748142619"></a><a name="ul15994748142619"></a><ul id="ul15994748142619"><li>Temporary files generated during multipart download are stored in this directory. Therefore, ensure that the user who executes obsutil has the write permission on the path.</li><li>The available space of the partition where the path is located must be greater than the size of the objects to be downloaded.</li></ul>
</div></div>
</td>
</tr>
<tr id="row1241495713314"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p156631810914"><a name="p156631810914"></a><a name="p156631810914"></a>dryRun</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p10663610617"><a name="p10663610617"></a><a name="p10663610617"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p76630101412"><a name="p76630101412"></a><a name="p76630101412"></a>Conducts a dry run.</p>
</td>
</tr>
<tr id="row35998457538"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p206001459530"><a name="p206001459530"></a><a name="p206001459530"></a>u</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p10601174595317"><a name="p10601174595317"></a><a name="p10601174595317"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1260204518534"><a name="p1260204518534"></a><a name="p1260204518534"></a>Indicates incremental download. If this parameter is set, each object can be downloaded only when it does not exist in the local path, its size is different from the namesake one in the local path, or it has the latest modification time.</p>
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
<div class="note" id="note17453113313289"><a name="note17453113313289"></a><a name="note17453113313289"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p12303311112719"><a name="p12303311112719"></a><a name="p12303311112719"></a>Objects in the bucket must contain metadata <strong id="b262077113412"><a name="b262077113412"></a><a name="b262077113412"></a>x-obs-md5chksum</strong>. Otherwise, MD5 verification will be skipped.</p>
</div></div>
</td>
</tr>
<tr id="row36131445115310"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p14613194514535"><a name="p14613194514535"></a><a name="p14613194514535"></a>p</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p15615164515317"><a name="p15615164515317"></a><a name="p15615164515317"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1161744513534"><a name="p1161744513534"></a><a name="p1161744513534"></a>Indicates the maximum number of concurrent multipart download tasks when downloading an object. The default value is the value of <strong id="b636105782416"><a name="b636105782416"></a><a name="b636105782416"></a>defaultParallels</strong> in the configuration file.</p>
</td>
</tr>
<tr id="row14617144520537"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p17619545185314"><a name="p17619545185314"></a><a name="p17619545185314"></a>threshold</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p9620645185311"><a name="p9620645185311"></a><a name="p9620645185311"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p262174575315"><a name="p262174575315"></a><a name="p262174575315"></a>Indicates the threshold for enabling multipart download, in bytes. The default value is the value of <strong id="b19309194493311"><a name="b19309194493311"></a><a name="b19309194493311"></a>defaultBigfileThreshold</strong> in the configuration file.</p>
<div class="note" id="note94444118613"><a name="note94444118613"></a><a name="note94444118613"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul182059695115"></a><a name="ul182059695115"></a><ul id="ul182059695115"><li>If the size of the object to be downloaded is smaller than the threshold, download the object directly. If not, a multipart download is required.</li><li>If you download an object directly, no part record is generated, and resumable transmission is not supported.</li><li>This value can contain a capacity unit. For example, <strong id="b119361641103410"><a name="b119361641103410"></a><a name="b119361641103410"></a>1 MB</strong> indicates 1048576 bytes.</li></ul>
</div></div>
</td>
</tr>
<tr id="row1363634595315"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p3637164545318"><a name="p3637164545318"></a><a name="p3637164545318"></a>ps</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p1264164517532"><a name="p1264164517532"></a><a name="p1264164517532"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p10641184516531"><a name="p10641184516531"></a><a name="p10641184516531"></a>Indicates the size of each part in a multipart download task, in bytes. The default value is the value of <strong id="b345384918346"><a name="b345384918346"></a><a name="b345384918346"></a>defaultPartSize</strong> in the configuration file.</p>
<div class="note" id="note1945367164418"><a name="note1945367164418"></a><a name="note1945367164418"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul148311025104617"></a><a name="ul148311025104617"></a><ul id="ul148311025104617"><li>This value can contain a capacity unit. For example, <strong id="b19867161105312"><a name="b19867161105312"></a><a name="b19867161105312"></a>1 MB</strong> indicates 1048576 bytes.</li><li>The parameter can be set to <strong id="b73321536533"><a name="b73321536533"></a><a name="b73321536533"></a>auto</strong>. In this case, obsutil automatically sets the part size for each multipart task based on the source object size.</li></ul>
</div></div>
</td>
</tr>
<tr id="row196461145155314"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p1764874511530"><a name="p1764874511530"></a><a name="p1764874511530"></a>cpd</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p206482455537"><a name="p206482455537"></a><a name="p206482455537"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1965018458536"><a name="p1965018458536"></a><a name="p1965018458536"></a>Indicates the folder where the part records reside. The default value is <strong id="b03101221102112"><a name="b03101221102112"></a><a name="b03101221102112"></a>.obsutil_checkpoint</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note519835512711"><a name="note519835512711"></a><a name="note519835512711"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p118866577266"><a name="p118866577266"></a><a name="p118866577266"></a>A part record is generated during a multipart download and saved to the <strong id="b17221185615478"><a name="b17221185615478"></a><a name="b17221185615478"></a>down</strong> subfolder. After the download succeeds, its part record is deleted automatically. If the download fails or is suspended, the system attempts to resume the task according to its part record when you perform the download the next time.</p>
</div></div>
</td>
</tr>
<tr id="row5655184520536"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p56577450535"><a name="p56577450535"></a><a name="p56577450535"></a>f</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p15712102082517"><a name="p15712102082517"></a><a name="p15712102082517"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p566014515533"><a name="p566014515533"></a><a name="p566014515533"></a>Runs in force mode.</p>
</td>
</tr>
<tr id="row1566014510533"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p3662204585313"><a name="p3662204585313"></a><a name="p3662204585313"></a>j</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p4719102092513"><a name="p4719102092513"></a><a name="p4719102092513"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p17665114555311"><a name="p17665114555311"></a><a name="p17665114555311"></a>Indicates the maximum number of concurrent tasks for downloading objects in batches. The default value is the value of <strong id="b862721920247"><a name="b862721920247"></a><a name="b862721920247"></a>defaultJobs</strong> in the configuration file.</p>
<div class="note" id="note891964620819"><a name="note891964620819"></a><a name="note891964620819"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p1091964618820"><a name="p1091964618820"></a><a name="p1091964618820"></a>The value is ensured to be greater than or equal to 1.</p>
</div></div>
</td>
</tr>
<tr id="row109762261097"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p156173113915"><a name="p156173113915"></a><a name="p156173113915"></a>exclude</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p1372419203256"><a name="p1372419203256"></a><a name="p1372419203256"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p16666614131519"><a name="p16666614131519"></a><a name="p16666614131519"></a>Indicates the matching patterns of source objects that are excluded, for example: <strong id="b2835104318103"><a name="b2835104318103"></a><a name="b2835104318103"></a>*.txt</strong>.</p>
<div class="note" id="note145284716208"><a name="note145284716208"></a><a name="note145284716208"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul05604713204"></a><a name="ul05604713204"></a><ul id="ul05604713204"><li>The asterisk (*) represents any group of characters, and the question mark (?) represents any single character. For instance, <strong id="b17220124561010"><a name="b17220124561010"></a><a name="b17220124561010"></a>abc*.txt</strong> indicates any file whose name starts with <strong id="b1922154518107"><a name="b1922154518107"></a><a name="b1922154518107"></a>abc</strong> and ends with <strong id="b1122313450103"><a name="b1122313450103"></a><a name="b1122313450103"></a>.txt</strong>.</li><li>You can use <strong id="b1734764681019"><a name="b1734764681019"></a><a name="b1734764681019"></a>\*</strong> to represent <strong id="b834814619103"><a name="b834814619103"></a><a name="b834814619103"></a>*</strong> and <strong id="b19349646101016"><a name="b19349646101016"></a><a name="b19349646101016"></a>\?</strong> to represent <strong id="b83501467102"><a name="b83501467102"></a><a name="b83501467102"></a>?</strong>.</li><li>If the name of the object to be downloaded matches the value of this parameter, the object is skipped.</li></ul>
</div></div>
<div class="notice" id="note179117549207"><a name="note179117549207"></a><a name="note179117549207"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul877892116516"></a><a name="ul877892116516"></a><ul id="ul877892116516"><li>You are advised to use quotation marks for the matching pattern to prevent special characters from being escaped by the OS and leading to unexpected results. Use single quotation marks for Linux or macOS and quotation marks for Windows.</li><li>The matching pattern applies to the absolute path of an object, including the object name prefix and object name starting from the root directory. For example, if the path of an object in the bucket is <strong id="b36591141163013"><a name="b36591141163013"></a><a name="b36591141163013"></a>obs://bucket/src1/src2/test.txt</strong>, then the absolute path of the object is <strong id="b66591141183017"><a name="b66591141183017"></a><a name="b66591141183017"></a>src1/src2/test.txt</strong>.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li></ul>
</div></div>
</td>
</tr>
<tr id="row20666154513530"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p186661845145315"><a name="p186661845145315"></a><a name="p186661845145315"></a>include</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p127271520102517"><a name="p127271520102517"></a><a name="p127271520102517"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p37071730153119"><a name="p37071730153119"></a><a name="p37071730153119"></a>Indicates the matching patterns of source objects that are included, for example: <strong id="b20113191471120"><a name="b20113191471120"></a><a name="b20113191471120"></a>*.jpg</strong>.</p>
<div class="note" id="note195168716220"><a name="note195168716220"></a><a name="note195168716220"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul752013715229"></a><a name="ul752013715229"></a><ul id="ul752013715229"><li>The asterisk (*) represents any group of characters, and the question mark (?) represents any single character.</li><li>You can use <strong id="b422113173113"><a name="b422113173113"></a><a name="b422113173113"></a>\*</strong> to represent <strong id="b102221617101111"><a name="b102221617101111"></a><a name="b102221617101111"></a>*</strong> and <strong id="b202231817151112"><a name="b202231817151112"></a><a name="b202231817151112"></a>\?</strong> to represent <strong id="b11224117141111"><a name="b11224117141111"></a><a name="b11224117141111"></a>?</strong>.</li><li>Only after identifying that the name of the file to be downloaded does not match the value of <strong id="b1874141931111"><a name="b1874141931111"></a><a name="b1874141931111"></a>exclude</strong>, the system checks whether the file name matches the value of this parameter. If yes, the file is downloaded. If not, the file is skipped.</li></ul>
</div></div>
<div class="notice" id="note9270217202212"><a name="note9270217202212"></a><a name="note9270217202212"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul12136162112586"></a><a name="ul12136162112586"></a><ul id="ul12136162112586"><li>You are advised to use quotation marks for the matching pattern to prevent special characters from being escaped by the OS and leading to unexpected results. Use single quotation marks for Linux or macOS and quotation marks for Windows.</li><li>The matching pattern applies to the absolute path of an object, including the object name prefix and object name starting from the root directory. For example, if the path of an object in the bucket is <strong id="b19602164573017"><a name="b19602164573017"></a><a name="b19602164573017"></a>obs://bucket/src1/src2/test.txt</strong>, then the absolute path of the object is <strong id="b1060314520305"><a name="b1060314520305"></a><a name="b1060314520305"></a>src1/src2/test.txt</strong>.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li></ul>
</div></div>
</td>
</tr>
<tr id="row64713716436"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p496314014481"><a name="p496314014481"></a><a name="p496314014481"></a>timeRange</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p1664510223494"><a name="p1664510223494"></a><a name="p1664510223494"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p16963940134813"><a name="p16963940134813"></a><a name="p16963940134813"></a>Indicates the time range matching pattern when downloading objects. Only objects whose last modification time is within the configured time range are downloaded.</p>
<p id="p15975633125310"><a name="p15975633125310"></a><a name="p15975633125310"></a>This pattern has a lower priority than the object matching patterns (<strong id="b8524185317313"><a name="b8524185317313"></a><a name="b8524185317313"></a>exclude</strong>/<strong id="b85251153539"><a name="b85251153539"></a><a name="b85251153539"></a>include</strong>). That is, the time range matching pattern is executed after the configured object matching patterns.</p>
<div class="note" id="note8766915165919"><a name="note8766915165919"></a><a name="note8766915165919"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul19272430145915"></a><a name="ul19272430145915"></a><ul id="ul19272430145915"><li>The matching time range is represented in <em id="i1297864783014"><a name="i1297864783014"></a><a name="i1297864783014"></a>time1</em><strong id="b997934713304"><a name="b997934713304"></a><a name="b997934713304"></a>-</strong><em id="i298074703013"><a name="i298074703013"></a><a name="i298074703013"></a>time2</em>, where <em id="i109803471301"><a name="i109803471301"></a><a name="i109803471301"></a>time1</em> must be earlier than or the same as <em id="i29811847203013"><a name="i29811847203013"></a><a name="i29811847203013"></a>time2</em>. The time format is <em id="i3982547153012"><a name="i3982547153012"></a><a name="i3982547153012"></a>yyyyMMddHHmmss</em>.</li><li>Automatic formatting is supported. For example, yyyyMMdd is equivalent to yyyyMMdd000000, and yyyyMM is equivalent to yyyyMM01000000.</li><li>If this parameter is set to <strong id="b5889165023014"><a name="b5889165023014"></a><a name="b5889165023014"></a>*-</strong><em id="i2890650113017"><a name="i2890650113017"></a><a name="i2890650113017"></a>time2</em>, all files whose last modification time is earlier than <em id="i68911450133018"><a name="i68911450133018"></a><a name="i68911450133018"></a>time2</em> are matched. If it is set to <em id="i489205019306"><a name="i489205019306"></a><a name="i489205019306"></a>time1</em><strong id="b1089215093014"><a name="b1089215093014"></a><a name="b1089215093014"></a>-*</strong>, all files whose last modification time is later than <em id="i98935500303"><a name="i98935500303"></a><a name="i98935500303"></a>time1</em> are matched.</li></ul>
</div></div>
<div class="notice" id="note1078981817591"><a name="note1078981817591"></a><a name="note1078981817591"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul881073612597"></a><a name="ul881073612597"></a><ul id="ul881073612597"><li>Time in the matching pattern is the UTC time.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li></ul>
</div></div>
</td>
</tr>
<tr id="row9683191411010"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p10122175516"><a name="p10122175516"></a><a name="p10122175516"></a>mf</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p1562018121550"><a name="p1562018121550"></a><a name="p1562018121550"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p31226718512"><a name="p31226718512"></a><a name="p31226718512"></a>Indicates that the name matching pattern (<strong id="b84215199612"><a name="b84215199612"></a><a name="b84215199612"></a>include</strong> or <strong id="b14210197613"><a name="b14210197613"></a><a name="b14210197613"></a>exclude</strong>) and the time matching pattern (<strong id="b194318191866"><a name="b194318191866"></a><a name="b194318191866"></a>timeRange</strong>) also take effect on objects whose names end with a slash (/).</p>
<p id="p15797105118514"><a name="p15797105118514"></a><a name="p15797105118514"></a></p>
</td>
</tr>
<tr id="row109093549288"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p143226572287"><a name="p143226572287"></a><a name="p143226572287"></a>o</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p3729820102512"><a name="p3729820102512"></a><a name="p3729820102512"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1632615718282"><a name="p1632615718282"></a><a name="p1632615718282"></a>Indicates the folder where operation result lists reside. After the command is executed, result lists (possibly including success, failure, and warning files) are generated in the folder. The default value is <strong id="b3998193418218"><a name="b3998193418218"></a><a name="b3998193418218"></a>.obsutil_output</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note1416817409247"><a name="note1416817409247"></a><a name="note1416817409247"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul101190347408"></a><a name="ul101190347408"></a><ul id="ul101190347408"><li>The naming rule for result lists is as follows: <strong id="b162206244343"><a name="b162206244343"></a><a name="b162206244343"></a>cp_{succeed  | failed | warning}_report_</strong><em id="i182214245342"><a name="i182214245342"></a><a name="i182214245342"></a>time</em><strong id="b14221142473413"><a name="b14221142473413"></a><a name="b14221142473413"></a>_TaskId.txt</strong><p id="li412103664016p0"><a name="li412103664016p0"></a><a name="li412103664016p0"></a>By default, the maximum size of a single result list is 30 MB and the maximum number of result lists that can be retained is 1024. You can set the maximum size and number by configuring <strong id="b07917272390"><a name="b07917272390"></a><a name="b07917272390"></a>recordMaxLogSize</strong> and <strong id="b67992716393"><a name="b67992716393"></a><a name="b67992716393"></a>recordBackups</strong> in the configuration file.</p>
</li></ul>
</div></div>
</td>
</tr>
<tr id="row27141035103619"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
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

-   Take the Windows OS as an example. Run the  **obsutil cp -recover=3066a4b0-4d21-4929-bb84-4829c32cbd0f d:\\ -f -r**  command to download objects in batches.

```
obsutil cp -recover=3066a4b0-4d21-4929-bb84-4829c32cbd0f -f -r

Parallel:      3                   Jobs:          3
Threshold:     524288000           PartSize:      5242880
Exclude:                           Include:
VerifyLength:  false               VerifyMd5:     false
CheckpointDir: xxxx
OutputDir: xxxx

[======================================================] 100.00% 155.59 KB/s 0s
Succeed count is:   1         Failed count is:    0
Metrics [max cost:153 ms, min cost:129 ms, average cost:92.00 ms, average tps:17.86]
Task id is: 19ad99ce-434e-41b2-9c8d-3af5b42eb65a
```

