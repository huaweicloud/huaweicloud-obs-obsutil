# Deleting Part Records<a name="EN-US_TOPIC_0141181370"></a>

## Function<a name="section1479112110815"></a>

You can use this command to delete part records from a specified directory.

## Command Line Structure<a name="section1220752192216"></a>

-   In Windows

    ```
    obsutil clear [checkpoint_dir] [-u] [-d] [-c] [-config=xxx]
    ```

-   In Linux or macOS

    ```
    ./obsutil clear [checkpoint_dir] [-u] [-d] [-c] [-config=xxx]
    ```


## Parameter Description<a name="section6559191102418"></a>

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="17%" id="mcps1.1.4.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="22%" id="mcps1.1.4.1.2"><p id="p15137125919108"><a name="p15137125919108"></a><a name="p15137125919108"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="61%" id="mcps1.1.4.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row108328217449"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p64495172515"><a name="p64495172515"></a><a name="p64495172515"></a>checkpoint_dir</p>
</td>
<td class="cellrowborder" valign="top" width="22%" headers="mcps1.1.4.1.2 "><p id="p154316502519"><a name="p154316502519"></a><a name="p154316502519"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="61%" headers="mcps1.1.4.1.3 "><p id="p17425512259"><a name="p17425512259"></a><a name="p17425512259"></a>Indicates the folder where the part records reside. The default value is <strong id="b5631204151111"><a name="b5631204151111"></a><a name="b5631204151111"></a>.obsutil_checkpoint</strong>, the same subfolder where obsutil commands reside.</p>
</td>
</tr>
<tr id="row32014962815"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p82054920283"><a name="p82054920283"></a><a name="p82054920283"></a>u</p>
</td>
<td class="cellrowborder" valign="top" width="22%" headers="mcps1.1.4.1.2 "><p id="p4454146102911"><a name="p4454146102911"></a><a name="p4454146102911"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="61%" headers="mcps1.1.4.1.3 "><p id="p52015497289"><a name="p52015497289"></a><a name="p52015497289"></a>Deletes the part records of all multipart upload tasks.</p>
<div class="note" id="note148764412348"><a name="note148764412348"></a><a name="note148764412348"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p184871144203418"><a name="p184871144203418"></a><a name="p184871144203418"></a>At the same time, the system attempts to delete the multipart upload tasks in the part records.</p>
</div></div>
</td>
</tr>
<tr id="row1155655010280"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p1655610501288"><a name="p1655610501288"></a><a name="p1655610501288"></a>d</p>
</td>
<td class="cellrowborder" valign="top" width="22%" headers="mcps1.1.4.1.2 "><p id="p124553652918"><a name="p124553652918"></a><a name="p124553652918"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="61%" headers="mcps1.1.4.1.3 "><p id="p1055655042814"><a name="p1055655042814"></a><a name="p1055655042814"></a>Deletes the part records of all multipart download tasks.</p>
<div class="note" id="note1018343543718"><a name="note1018343543718"></a><a name="note1018343543718"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p18183183523719"><a name="p18183183523719"></a><a name="p18183183523719"></a>At the same time, the system attempts to delete the fragments in the part records.</p>
</div></div>
</td>
</tr>
<tr id="row5616115192814"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p561675112281"><a name="p561675112281"></a><a name="p561675112281"></a>c</p>
</td>
<td class="cellrowborder" valign="top" width="22%" headers="mcps1.1.4.1.2 "><p id="p745815619297"><a name="p745815619297"></a><a name="p745815619297"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="61%" headers="mcps1.1.4.1.3 "><p id="p14616651172819"><a name="p14616651172819"></a><a name="p14616651172819"></a>Deletes the part records of all multipart copy tasks.</p>
<div class="note" id="note28911225113717"><a name="note28911225113717"></a><a name="note28911225113717"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p12892112516371"><a name="p12892112516371"></a><a name="p12892112516371"></a>At the same time, the system attempts to delete the multipart copy tasks in the part records.</p>
</div></div>
</td>
</tr>
<tr id="row20579102283818"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="22%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="61%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

>![](public_sys-resources/icon-note.gif) **NOTE:**   
>You must configure at least one among the  **u**,  **d**  and  **c**  parameters.  

## Running Example<a name="section15899161919244"></a>

-   Take the Windows OS as an example. Run the  **obsutil clear -u**  command to delete the part records of multipart upload tasks in the default directory.

```
obsutil clear -u

Clear checkpoint files for uploading in folder [xxxxx]
[==================================================================] 100.00% 0s
Succeed files is:   1         Failed files is:    0
```

