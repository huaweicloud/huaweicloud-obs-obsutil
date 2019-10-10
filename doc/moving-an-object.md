# Moving an Object<a name="EN-US_TOPIC_0164915882"></a>

## Function<a name="section2046320316257"></a>

You can use this command to move a single object or move objects in batches by a specified object name prefix.

>![](public_sys-resources/icon-notice.gif) **NOTICE:**   
>-   Do not change the source objects in the OBS bucket when moving objects. Otherwise, the operation may fail or data may be inconsistent.  
>-   The source objects are deleted after the move operation succeeds.  

## Command Line Structure<a name="section23858124111"></a>

-   In Windows
    -   Moving a single object

        ```
        obsutil mv obs://srcbucket/key obs://dstbucket/[dest] [-dryRun] [-u] [-p=1] [-threshold=52428800] [-versionId=xxx] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-cpd=xxx] [-fr] [-o=xxx] [-config=xxx]
        ```

    -   Moving objects in batches

        ```
        obsutil mv obs://srcbucket[/key] obs://dstbucket[/dest] -r [-dryRun] [-f] [-flat] [-u] [-j=1] [-p=1] [-threshold=52428800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]
        ```


-   In Linux OS or macOS
    -   Moving a single object

        ```
        ./obsutil mv obs://srcbucket/key obs://dstbucket/[dest] [-dryRun] [-u] [-p=1] [-threshold=52428800] [-versionId=xxx] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-cpd=xxx] [-fr] [-o=xxx] [-config=xxx]
        ```

    -   Moving objects in batches

        ```
        ./obsutil mv obs://srcbucket[/key] obs://dstbucket[/dest] -r [-dryRun] [-f] [-flat] [-u] [-j=1] [-p=1] [-threshold=52428800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]
        ```



>![](public_sys-resources/icon-note.gif) **NOTE:**   
>-   The source path and destination path cannot be the same.  
>-   The source and destination paths cannot be nested when moving objects in batches.  

## Parameter Description<a name="section33057574522"></a>

<a name="table125609307312"></a>
<table><thead align="left"><tr id="row456603043117"><th class="cellrowborder" valign="top" width="17.57%" id="mcps1.1.4.1.1"><p id="p85680303316"><a name="p85680303316"></a><a name="p85680303316"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="18.529999999999998%" id="mcps1.1.4.1.2"><p id="p14572130153112"><a name="p14572130153112"></a><a name="p14572130153112"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="63.9%" id="mcps1.1.4.1.3"><p id="p15751930133118"><a name="p15751930133118"></a><a name="p15751930133118"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row20577630203112"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p18579630143115"><a name="p18579630143115"></a><a name="p18579630143115"></a>srcbucket</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p458314302313"><a name="p458314302313"></a><a name="p458314302313"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p34544306544"><a name="p34544306544"></a><a name="p34544306544"></a>Source bucket name</p>
</td>
</tr>
<tr id="row1059316307314"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p159616306318"><a name="p159616306318"></a><a name="p159616306318"></a>dstbucket</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p1559933073119"><a name="p1559933073119"></a><a name="p1559933073119"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p95994308315"><a name="p95994308315"></a><a name="p95994308315"></a>Destination bucket name</p>
</td>
</tr>
<tr id="row169347719447"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p1811911004419"><a name="p1811911004419"></a><a name="p1811911004419"></a>dest</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p9123510204415"><a name="p9123510204415"></a><a name="p9123510204415"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p1839313263477"><a name="p1839313263477"></a><a name="p1839313263477"></a>Indicates the destination object name when moving a single object, or the name prefix of destination objects when moving objects in batches.</p>
</td>
</tr>
<tr id="row106007303311"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p1060363010318"><a name="p1060363010318"></a><a name="p1060363010318"></a>key</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p8535551165513"><a name="p8535551165513"></a><a name="p8535551165513"></a>Mandatory for moving a single object</p>
<p id="p7265911561"><a name="p7265911561"></a><a name="p7265911561"></a>Optional for moving objects in batches</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p176052030183114"><a name="p176052030183114"></a><a name="p176052030183114"></a>Indicates the source object name when moving a single object, or the name prefix of source objects when moving objects in batches.</p>
<p id="p36061309319"><a name="p36061309319"></a><a name="p36061309319"></a>The rules are as follows:</p>
<a name="ul1560743053116"></a><a name="ul1560743053116"></a><ul id="ul1560743053116"><li>This parameter cannot be left blank when moving a single object. If <strong id="b17382225533"><a name="b17382225533"></a><a name="b17382225533"></a>dest</strong> is left blank, the source object is moved to the root directory of the destination bucket. If the value of <strong id="b838316257318"><a name="b838316257318"></a><a name="b838316257318"></a>dest</strong> ends with a slash (/), the destination object name is the value of <strong id="b238462516311"><a name="b238462516311"></a><a name="b238462516311"></a>dest</strong> plus the source object name. Otherwise, the destination object name is the value of <strong id="b14384325336"><a name="b14384325336"></a><a name="b14384325336"></a>dest</strong>.</li><li>If this parameter is left blank when moving objects in batches, all objects in the source bucket are moved. If not, objects whose name prefix is the set value in the source bucket are moved. The rules for confirming the name of the destination object are as follows:<a name="ul101911711114619"></a><a name="ul101911711114619"></a><ul id="ul101911711114619"><li>If the value of <strong id="b175610418258"><a name="b175610418258"></a><a name="b175610418258"></a>dest</strong> ends with a slash (/), the destination object name is the value of <strong id="b1756312281499"><a name="b1756312281499"></a><a name="b1756312281499"></a>dest</strong> plus the source object name.</li><li>If the value of <strong id="b198673121165"><a name="b198673121165"></a><a name="b198673121165"></a>dest</strong> does not end with a slash (/), the destination object name is <em id="i1144918511760"><a name="i1144918511760"></a><a name="i1144918511760"></a>the value of </em><em id="i651522503618"><a name="i651522503618"></a><a name="i651522503618"></a><strong id="b92851281163"><a name="b92851281163"></a><a name="b92851281163"></a>dest</strong></em><strong id="b10281031178"><a name="b10281031178"></a><a name="b10281031178"></a>/</strong><em id="i563313811717"><a name="i563313811717"></a><a name="i563313811717"></a>source object name</em>.</li></ul>
</li></ul>
<div class="note" id="note14528254193716"><a name="note14528254193716"></a><a name="note14528254193716"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul17203733184419"></a><a name="ul17203733184419"></a><ul id="ul17203733184419"><li>If this parameter is configured but parameter <strong id="b882713482518"><a name="b882713482518"></a><a name="b882713482518"></a>flat</strong> is not when moving objects in batches, the name of the source object contains the name prefix of the parent object. If <strong id="b168287489512"><a name="b168287489512"></a><a name="b168287489512"></a>flat</strong> is configured, then the name of the source object does not contain the name prefix of the parent object.</li><li>For details about how to use this parameter, see <a href="#section23858124111">Command Line Structure</a>.</li></ul>
</div></div>
</td>
</tr>
<tr id="row64161323113"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p123051451513"><a name="p123051451513"></a><a name="p123051451513"></a>fr</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p113071451811"><a name="p113071451811"></a><a name="p113071451811"></a>Optional for moving an object (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p73101351816"><a name="p73101351816"></a><a name="p73101351816"></a>Generates an operation result list when moving an object.</p>
</td>
</tr>
<tr id="row19803114816115"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p13236151151117"><a name="p13236151151117"></a><a name="p13236151151117"></a>flat</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p142391351111116"><a name="p142391351111116"></a><a name="p142391351111116"></a>Optional for moving objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p624025171110"><a name="p624025171110"></a><a name="p624025171110"></a>The name prefix of the parent object is excluded when moving objects in batches.</p>
</td>
</tr>
<tr id="row1513173918563"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p1537517155414"><a name="p1537517155414"></a><a name="p1537517155414"></a>dryRun</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p13376216548"><a name="p13376216548"></a><a name="p13376216548"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p1137981175411"><a name="p1137981175411"></a><a name="p1137981175411"></a>Conducts a dry run.</p>
</td>
</tr>
<tr id="row17612113018317"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p1461418302318"><a name="p1461418302318"></a><a name="p1461418302318"></a>u</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p1615183013110"><a name="p1615183013110"></a><a name="p1615183013110"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p12647141718263"><a name="p12647141718263"></a><a name="p12647141718263"></a>Indicates incremental move. If this parameter is set, each object can be moved only when it does not exist in the destination bucket, its size is different from the namesake one in the destination bucket, or it has the latest modification time.</p>
<div class="note" id="note21654238262"><a name="note21654238262"></a><a name="note21654238262"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p756433614817"><a name="p756433614817"></a><a name="p756433614817"></a>If the size and modification time of the destination object are the same as those of the source object, the source object is directly deleted instead of being moved.</p>
</div></div>
</td>
</tr>
<tr id="row3632193093117"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p26341030163117"><a name="p26341030163117"></a><a name="p26341030163117"></a>p</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p9636183093120"><a name="p9636183093120"></a><a name="p9636183093120"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p7638163013112"><a name="p7638163013112"></a><a name="p7638163013112"></a>Indicates the maximum number of concurrent multipart move tasks when moving an object. The default value is the value of <strong id="b495133214444"><a name="b495133214444"></a><a name="b495133214444"></a>defaultParallels</strong> in the configuration file.</p>
</td>
</tr>
<tr id="row363913014313"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p116401030183113"><a name="p116401030183113"></a><a name="p116401030183113"></a>threshold</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p176431304319"><a name="p176431304319"></a><a name="p176431304319"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p16441830113114"><a name="p16441830113114"></a><a name="p16441830113114"></a>Indicates the threshold for enabling multipart move, in bytes. The default value is the value of <strong id="b289264273012"><a name="b289264273012"></a><a name="b289264273012"></a>defaultBigfileThreshold</strong> in the configuration file.</p>
<div class="note" id="note1382352674913"><a name="note1382352674913"></a><a name="note1382352674913"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul182059695115"></a><a name="ul182059695115"></a><ul id="ul182059695115"><li>If the size of the object to be moved is smaller than the threshold, move the object directly. If not, a multipart move is required.</li><li>If you move an object directly, no part record is generated, and resumable transmission is not supported.</li><li>This value can contain a capacity unit. For example, <strong id="b114201348164"><a name="b114201348164"></a><a name="b114201348164"></a>1 MB</strong> indicates 1048576 bytes.</li></ul>
</div></div>
</td>
</tr>
<tr id="row145994924514"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p95911457154514"><a name="p95911457154514"></a><a name="p95911457154514"></a>versionId</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p0594135774510"><a name="p0594135774510"></a><a name="p0594135774510"></a>Optional for moving an object (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p1859655744514"><a name="p1859655744514"></a><a name="p1859655744514"></a>Source object version ID that can be specified when moving a single object</p>
</td>
</tr>
<tr id="row264453012315"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p964614304315"><a name="p964614304315"></a><a name="p964614304315"></a>acl</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p17648193011318"><a name="p17648193011318"></a><a name="p17648193011318"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p16650163083117"><a name="p16650163083117"></a><a name="p16650163083117"></a>Access control policies for destination objects that can be specified when moving objects. Possible values are:</p>
<a name="ul0651183053115"></a><a name="ul0651183053115"></a><ul id="ul0651183053115"><li>private</li><li>public-read</li><li>public-read-write</li><li>bucket-owner-full-control</li></ul>
<div class="note" id="note1790113183525"><a name="note1790113183525"></a><a name="note1790113183525"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p92982882916"><a name="p92982882916"></a><a name="p92982882916"></a>The preceding four values indicate private read and write, public read, public read and write, and bucket owner full control.</p>
</div></div>
</td>
</tr>
<tr id="row62441623105116"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p13466172515511"><a name="p13466172515511"></a><a name="p13466172515511"></a>sc</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p134681525175117"><a name="p134681525175117"></a><a name="p134681525175117"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p04731025135114"><a name="p04731025135114"></a><a name="p04731025135114"></a>Storage classes of the destination objects that can be specified when moving objects. Possible values are:</p>
<a name="ul175651814214"></a><a name="ul175651814214"></a><ul id="ul175651814214"><li><strong id="b148691033134015"><a name="b148691033134015"></a><a name="b148691033134015"></a>standard</strong>: OBS Standard, which features low access latency and high throughput, and is applicable to storing frequently accessed data (multiple accesses per month averagely) or data that is smaller than 1 MB</li><li><strong id="b75593541842"><a name="b75593541842"></a><a name="b75593541842"></a>warm</strong>: OBS Warm. It is applicable to storing semi-frequently accessed (less than 12 times a year averagely) data that requires quick response.</li><li><strong id="b11827701056"><a name="b11827701056"></a><a name="b11827701056"></a>cold</strong>: OBS Cold. It is secure, durable, and inexpensive, and applicable to archiving rarely-accessed (once a year averagely) data.</li></ul>
</td>
</tr>
<tr id="row8659203023112"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p1366223093111"><a name="p1366223093111"></a><a name="p1366223093111"></a>meta</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p26631230193118"><a name="p26631230193118"></a><a name="p26631230193118"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p17941525135111"><a name="p17941525135111"></a><a name="p17941525135111"></a>Metadata of destination objects that can be specified when copying objects. The format is <em id="i183948142313"><a name="i183948142313"></a><a name="i183948142313"></a>key1</em><strong id="b839719147310"><a name="b839719147310"></a><a name="b839719147310"></a>:</strong><em id="i143970148316"><a name="i143970148316"></a><a name="i143970148316"></a>value1</em><strong id="b639819141631"><a name="b639819141631"></a><a name="b639819141631"></a>#</strong><em id="i2039812142317"><a name="i2039812142317"></a><a name="i2039812142317"></a>key2</em><strong id="b63995141316"><a name="b63995141316"></a><a name="b63995141316"></a>:</strong><em id="i173991141735"><a name="i173991141735"></a><a name="i173991141735"></a>value2</em><strong id="b940011141031"><a name="b940011141031"></a><a name="b940011141031"></a>#</strong><em id="i34011314436"><a name="i34011314436"></a><a name="i34011314436"></a>key3</em><strong id="b1340119143316"><a name="b1340119143316"></a><a name="b1340119143316"></a>:</strong><em id="i1940212141535"><a name="i1940212141535"></a><a name="i1940212141535"></a>value3</em>.</p>
<div class="note" id="note1606229165119"><a name="note1606229165119"></a><a name="note1606229165119"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p686342013559"><a name="p686342013559"></a><a name="p686342013559"></a>The preceding value indicates that the destination objects in the bucket contain three groups of customized metadata after objects are copied: <strong id="b17135731716"><a name="b17135731716"></a><a name="b17135731716"></a>key1:value1</strong>, <strong id="b613177161713"><a name="b613177161713"></a><a name="b613177161713"></a>key2:value2</strong>, and <strong id="b013167111711"><a name="b013167111711"></a><a name="b013167111711"></a>key3:value3</strong>.</p>
</div></div>
</td>
</tr>
<tr id="row466716303317"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p1366917307314"><a name="p1366917307314"></a><a name="p1366917307314"></a>ps</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p267203023115"><a name="p267203023115"></a><a name="p267203023115"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p1167313023117"><a name="p1167313023117"></a><a name="p1167313023117"></a>Indicates the size of each part in a multipart move task, in bytes. The value ranges from 100 KB to 5 GB. The default value is the value of <strong id="b1782672311313"><a name="b1782672311313"></a><a name="b1782672311313"></a>defaultPartSize</strong> in the configuration file.</p>
<div class="note" id="note12706112664318"><a name="note12706112664318"></a><a name="note12706112664318"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul67581919204219"></a><a name="ul67581919204219"></a><ul id="ul67581919204219"><li>This value can contain a capacity unit. For example, <strong id="b15780167439"><a name="b15780167439"></a><a name="b15780167439"></a>1 MB</strong> indicates 1048576 bytes.</li><li>The parameter can be set to <strong id="b8938151934315"><a name="b8938151934315"></a><a name="b8938151934315"></a>auto</strong>. In this case, obsutil automatically sets the part size for each multipart task based on the source object size.</li></ul>
</div></div>
</td>
</tr>
<tr id="row1367463093119"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p1967612305318"><a name="p1967612305318"></a><a name="p1967612305318"></a>cpd</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p1367810308316"><a name="p1367810308316"></a><a name="p1367810308316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p46792307316"><a name="p46792307316"></a><a name="p46792307316"></a>Indicates the folder where the part records reside. The default value is <strong id="b13201154217257"><a name="b13201154217257"></a><a name="b13201154217257"></a>.obsutil_checkpoint</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note22601842192718"><a name="note22601842192718"></a><a name="note22601842192718"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p118866577266"><a name="p118866577266"></a><a name="p118866577266"></a>A part record is generated during a multipart move and saved to the <strong id="b2079243114173"><a name="b2079243114173"></a><a name="b2079243114173"></a>copy</strong> subfolder. After the move succeeds, its part record is deleted automatically. If the move fails or is suspended, the system attempts to resume the task according to its part record when you perform the move the next time.</p>
</div></div>
</td>
</tr>
<tr id="row12681173043113"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p1683230123111"><a name="p1683230123111"></a><a name="p1683230123111"></a>r</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p1368433083110"><a name="p1368433083110"></a><a name="p1368433083110"></a>Mandatory for moving objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p12686930183119"><a name="p12686930183119"></a><a name="p12686930183119"></a>Moves objects in batches based on a specified name prefix of objects in the source bucket.</p>
</td>
</tr>
<tr id="row1568763019310"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p2688183018315"><a name="p2688183018315"></a><a name="p2688183018315"></a>f</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p1691193012317"><a name="p1691193012317"></a><a name="p1691193012317"></a>Optional for moving objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p1269315302318"><a name="p1269315302318"></a><a name="p1269315302318"></a>Runs in force mode.</p>
</td>
</tr>
<tr id="row20694163013111"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p269613053110"><a name="p269613053110"></a><a name="p269613053110"></a>j</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p170013055218"><a name="p170013055218"></a><a name="p170013055218"></a>Optional for moving objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p4699130153112"><a name="p4699130153112"></a><a name="p4699130153112"></a>Indicates the maximum number of concurrent tasks for moving objects in batches. The default value is the value of <strong id="b6802155075017"><a name="b6802155075017"></a><a name="b6802155075017"></a>defaultJobs</strong> in the configuration file.</p>
<div class="note" id="note891964620819"><a name="note891964620819"></a><a name="note891964620819"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p1091964618820"><a name="p1091964618820"></a><a name="p1091964618820"></a>The value is ensured to be greater than or equal to 1.</p>
</div></div>
</td>
</tr>
<tr id="row8334161131517"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p7660121421518"><a name="p7660121421518"></a><a name="p7660121421518"></a>exclude</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p176621714131512"><a name="p176621714131512"></a><a name="p176621714131512"></a>Optional for moving objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p16666614131519"><a name="p16666614131519"></a><a name="p16666614131519"></a>Indicates the matching patterns of source objects that are excluded, for example: <strong id="b98036164282"><a name="b98036164282"></a><a name="b98036164282"></a>*.txt</strong>.</p>
<div class="note" id="note145284716208"><a name="note145284716208"></a><a name="note145284716208"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul05604713204"></a><a name="ul05604713204"></a><ul id="ul05604713204"><li>The asterisk (*) represents any group of characters, and the question mark (?) represents any single character. For instance, <strong id="b220217189286"><a name="b220217189286"></a><a name="b220217189286"></a>abc*.txt</strong> indicates any file whose name starts with <strong id="b192030188286"><a name="b192030188286"></a><a name="b192030188286"></a>abc</strong> and ends with <strong id="b19205318162818"><a name="b19205318162818"></a><a name="b19205318162818"></a>.txt</strong>.</li><li>You can use <strong id="b95182011288"><a name="b95182011288"></a><a name="b95182011288"></a>\*</strong> to represent <strong id="b86520152817"><a name="b86520152817"></a><a name="b86520152817"></a>*</strong> and <strong id="b147122092814"><a name="b147122092814"></a><a name="b147122092814"></a>\?</strong> to represent <strong id="b168122032817"><a name="b168122032817"></a><a name="b168122032817"></a>?</strong>.</li><li>If the name of the object to be moved matches the value of this parameter, the object is skipped.</li></ul>
</div></div>
<div class="notice" id="note179117549207"><a name="note179117549207"></a><a name="note179117549207"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul877892116516"></a><a name="ul877892116516"></a><ul id="ul877892116516"><li>You are advised to use quotation marks for the matching pattern to prevent special characters from being escaped by the OS and leading to unexpected results. Use single quotation marks for Linux or macOS and quotation marks for Windows.</li><li>The matching pattern applies to the absolute path of an object, including the object name prefix and object name starting from the root directory. For example, if the path of an object in the bucket is <strong id="b17743109155712"><a name="b17743109155712"></a><a name="b17743109155712"></a>obs://bucket/src1/src2/test.txt</strong>, then the absolute path of the object is <strong id="b0744995574"><a name="b0744995574"></a><a name="b0744995574"></a>src1/src2/test.txt</strong>.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li></ul>
</div></div>
</td>
</tr>
<tr id="row207003306314"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p11703123012317"><a name="p11703123012317"></a><a name="p11703123012317"></a>include</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p470523013527"><a name="p470523013527"></a><a name="p470523013527"></a>Optional for moving objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p37071730153119"><a name="p37071730153119"></a><a name="p37071730153119"></a>Indicates the matching patterns of source objects that are included, for example: <strong id="b17984162514281"><a name="b17984162514281"></a><a name="b17984162514281"></a>*.jpg</strong>.</p>
<div class="note" id="note195168716220"><a name="note195168716220"></a><a name="note195168716220"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul752013715229"></a><a name="ul752013715229"></a><ul id="ul752013715229"><li>The asterisk (*) represents any group of characters, and the question mark (?) represents any single character.</li><li>You can use <strong id="b144924521579"><a name="b144924521579"></a><a name="b144924521579"></a>\*</strong> to represent <strong id="b1349314526578"><a name="b1349314526578"></a><a name="b1349314526578"></a>*</strong> and <strong id="b649414527570"><a name="b649414527570"></a><a name="b649414527570"></a>\?</strong> to represent <strong id="b184951752155711"><a name="b184951752155711"></a><a name="b184951752155711"></a>?</strong>.</li><li>Only after identifying that the name of the file to be moved does not match the value of <strong id="b186317308287"><a name="b186317308287"></a><a name="b186317308287"></a>exclude</strong>, the system checks whether the file name matches the value of this parameter. If yes, the file is moved. If not, the file is skipped.</li></ul>
</div></div>
<div class="notice" id="note9270217202212"><a name="note9270217202212"></a><a name="note9270217202212"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul11495105419578"></a><a name="ul11495105419578"></a><ul id="ul11495105419578"><li>You are advised to use quotation marks for the matching pattern to prevent special characters from being escaped by the OS and leading to unexpected results. Use single quotation marks for Linux or macOS and quotation marks for Windows.</li><li>The matching pattern applies to the absolute path of an object, including the object name prefix and object name starting from the root directory. For example, if the path of an object in the bucket is <strong id="b19761653185911"><a name="b19761653185911"></a><a name="b19761653185911"></a>obs://bucket/src1/src2/test.txt</strong>, then the absolute path of the object is <strong id="b1697635319597"><a name="b1697635319597"></a><a name="b1697635319597"></a>src1/src2/test.txt</strong>.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li></ul>
</div></div>
</td>
</tr>
<tr id="row881016184213"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p496314014481"><a name="p496314014481"></a><a name="p496314014481"></a>timeRange</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p1664510223494"><a name="p1664510223494"></a><a name="p1664510223494"></a>Optional for moving objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p16963940134813"><a name="p16963940134813"></a><a name="p16963940134813"></a>Indicates the time range matching pattern when moving objects. Only objects whose last modification time is within the configured time range are moved.</p>
<p id="p1565821819529"><a name="p1565821819529"></a><a name="p1565821819529"></a>This pattern has a lower priority than the object matching patterns (<strong id="b798717286019"><a name="b798717286019"></a><a name="b798717286019"></a>exclude</strong>/<strong id="b898918288012"><a name="b898918288012"></a><a name="b898918288012"></a>include</strong>). That is, the time range matching pattern is executed after the configured object matching patterns.</p>
<div class="note" id="note8766915165919"><a name="note8766915165919"></a><a name="note8766915165919"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul19272430145915"></a><a name="ul19272430145915"></a><ul id="ul19272430145915"><li>The matching time range is represented in <em id="i166591939609"><a name="i166591939609"></a><a name="i166591939609"></a>time1</em><strong id="b146611839208"><a name="b146611839208"></a><a name="b146611839208"></a>-</strong><em id="i196626391809"><a name="i196626391809"></a><a name="i196626391809"></a>time2</em>, where <em id="i2663739401"><a name="i2663739401"></a><a name="i2663739401"></a>time1</em> must be earlier than or the same as <em id="i866513913015"><a name="i866513913015"></a><a name="i866513913015"></a>time2</em>. The time format is <em id="i3666639806"><a name="i3666639806"></a><a name="i3666639806"></a>yyyyMMddHHmmss</em>.</li><li>Automatic formatting is supported. For example, yyyyMMdd is equivalent to yyyyMMdd000000, and yyyyMM is equivalent to yyyyMM01000000.</li><li>If this parameter is set to <strong id="b139130461105"><a name="b139130461105"></a><a name="b139130461105"></a>*-</strong><em id="i18914846808"><a name="i18914846808"></a><a name="i18914846808"></a>time2</em>, all files whose last modification time is earlier than <em id="i491514461408"><a name="i491514461408"></a><a name="i491514461408"></a>time2</em> are matched. If it is set to <em id="i189172463010"><a name="i189172463010"></a><a name="i189172463010"></a>time1</em><strong id="b291810461016"><a name="b291810461016"></a><a name="b291810461016"></a>-*</strong>, all files whose last modification time is later than <em id="i17919104617014"><a name="i17919104617014"></a><a name="i17919104617014"></a>time1</em> are matched.</li></ul>
</div></div>
<div class="notice" id="note1078981817591"><a name="note1078981817591"></a><a name="note1078981817591"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul881073612597"></a><a name="ul881073612597"></a><ul id="ul881073612597"><li>Time in the matching pattern is the UTC time.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li></ul>
</div></div>
</td>
</tr>
<tr id="row17951173311713"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p10122175516"><a name="p10122175516"></a><a name="p10122175516"></a>mf</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p1562018121550"><a name="p1562018121550"></a><a name="p1562018121550"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p31226718512"><a name="p31226718512"></a><a name="p31226718512"></a>Indicates that the name matching pattern (<strong id="b1218316342316"><a name="b1218316342316"></a><a name="b1218316342316"></a>include</strong> or <strong id="b5196634936"><a name="b5196634936"></a><a name="b5196634936"></a>exclude</strong>) and the time matching pattern (<strong id="b201971134335"><a name="b201971134335"></a><a name="b201971134335"></a>timeRange</strong>) also take effect on objects whose names end with a slash (/).</p>
<p id="p15797105118514"><a name="p15797105118514"></a><a name="p15797105118514"></a></p>
</td>
</tr>
<tr id="row9725193023119"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p1672753013119"><a name="p1672753013119"></a><a name="p1672753013119"></a>o</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p871493095216"><a name="p871493095216"></a><a name="p871493095216"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p137321030143112"><a name="p137321030143112"></a><a name="p137321030143112"></a>Indicates the folder where operation result lists reside. After the command is executed, result lists (possibly including success, failure, and warning files) are generated in the folder. The default value is <strong id="b68914575255"><a name="b68914575255"></a><a name="b68914575255"></a>.obsutil_output</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note289083220249"><a name="note289083220249"></a><a name="note289083220249"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul101190347408"></a><a name="ul101190347408"></a><ul id="ul101190347408"><li>The naming rule for result lists is as follows: <strong id="b184012103119"><a name="b184012103119"></a><a name="b184012103119"></a>mv_{succeed  | failed | warning}_report_</strong><em id="i1040215107118"><a name="i1040215107118"></a><a name="i1040215107118"></a>time</em><strong id="b11404131011118"><a name="b11404131011118"></a><a name="b11404131011118"></a>_TaskId.txt</strong></li><li>By default, the maximum size of a single result list is 30 MB and the maximum number of result lists that can be retained is 1024. You can set the maximum size and number by configuring <strong id="b8246121873610"><a name="b8246121873610"></a><a name="b8246121873610"></a>recordMaxLogSize</strong> and <strong id="b102471818103610"><a name="b102471818103610"></a><a name="b102471818103610"></a>recordBackups</strong> in the configuration file.</li></ul>
</div></div>
</td>
</tr>
<tr id="row1112831783618"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Response<a name="section1525012533192"></a>

Refer to  [Response](uploading-an-object.md#section6926520122416)  for uploading an object.

## Running Examples<a name="section15979161318211"></a>

-   Take the Windows OS as an example. Run the  **obsutil mv obs://bucket-test/key obs://bucket-test2**  command to move a single object.

```
obsutil mv obs://bucket-test/key obs://bucket-test2

Parallel:      3                   Jobs:          3
Threshold:     524288000           PartSize:      5242880
Exclude:                           Include:
VerifyLength:  false               VerifyMd5:     false
CheckpointDir: xxxx

[=====================================================] 100.00% 6/s 0s
Move successfully, 19B, obs://bucket-test/key --> obs://bucket-test2/key
ext.txt
```

-   Take the Windows OS as an example. Run the  **obsutil mv obs://bucket-test/temp/ obs://bucket-test2 -f -r**  command to move objects in batches.

```
obsutil mv obs://bucket-test/temp/ obs://bucket-test2 -f -r

Parallel:      3                   Jobs:          3
Threshold:     524288000           PartSize:      5242880
Exclude:                           Include:
VerifyLength:  false               VerifyMd5:     false
CheckpointDir: xxxx
OutputDir: xxxx


[=============================================================] 100.00% 10/s 0s
Succeed count is:   5         Failed count is:    0
Metrics [max cost:298 ms, min cost:192 ms, average cost:238.00 ms, average tps:9.71]
Task id is: 0476929d-9d23-4dc5-b2f8-0a0493f027c5
```

