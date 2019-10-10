# Updating a Configuration File<a name="EN-US_TOPIC_0141181369"></a>

## Function<a name="section1479112110815"></a>

You can use this command to update some configurations in the  **.obsutilconfig**  configuration file.

## Command Line Structure<a name="section1220752192216"></a>

-   In Windows

    ```
    obsutil config -interactive [-crr] [-config=xxx]
    ```

-   In Linux or macOS

    ```
    ./obsutil config -interactive [-crr] [-config=xxx]
    ```


## Parameter Description<a name="section6559191102418"></a>

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="16%" id="mcps1.1.4.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="28.999999999999996%" id="mcps1.1.4.1.2"><p id="p15137125919108"><a name="p15137125919108"></a><a name="p15137125919108"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="55.00000000000001%" id="mcps1.1.4.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row16853134519558"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p8853194565517"><a name="p8853194565517"></a><a name="p8853194565517"></a>interactive</p>
</td>
<td class="cellrowborder" valign="top" width="28.999999999999996%" headers="mcps1.1.4.1.2 "><p id="p128627553194"><a name="p128627553194"></a><a name="p128627553194"></a>Mandatory (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1585474515513"><a name="p1585474515513"></a><a name="p1585474515513"></a>Updates configurations in interactive mode.</p>
</td>
</tr>
<tr id="row95195143518"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p185285143520"><a name="p185285143520"></a><a name="p185285143520"></a>crr</p>
</td>
<td class="cellrowborder" valign="top" width="28.999999999999996%" headers="mcps1.1.4.1.2 "><p id="p466619556356"><a name="p466619556356"></a><a name="p466619556356"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1052135112358"><a name="p1052135112358"></a><a name="p1052135112358"></a>Updates the configurations related to client-side cross-region replication in the configuration file.</p>
<div class="note" id="note1841535164114"><a name="note1841535164114"></a><a name="note1841535164114"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p784163511412"><a name="p784163511412"></a><a name="p784163511412"></a>If this parameter is set, fields in the configuration file corresponding to parameters <strong id="b17483442435"><a name="b17483442435"></a><a name="b17483442435"></a>e</strong>, <strong id="b15867144612431"><a name="b15867144612431"></a><a name="b15867144612431"></a>i</strong>, <strong id="b27528508437"><a name="b27528508437"></a><a name="b27528508437"></a>k</strong>, and <strong id="b117531253174313"><a name="b117531253174313"></a><a name="b117531253174313"></a>t</strong> are respectively changed to <strong id="b13673175820438"><a name="b13673175820438"></a><a name="b13673175820438"></a>endpointCrr</strong>, <strong id="b11638219449"><a name="b11638219449"></a><a name="b11638219449"></a>akCrr</strong>, <strong id="b125108544411"><a name="b125108544411"></a><a name="b125108544411"></a>skCrr</strong>, and <strong id="b1014619108447"><a name="b1014619108447"></a><a name="b1014619108447"></a>tokenCrr</strong>.</p>
</div></div>
</td>
</tr>
<tr id="row9445101517387"><td class="cellrowborder" valign="top" width="16%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="28.999999999999996%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Running Example<a name="section15899161919244"></a>

-   Take the Windows OS as an example. Run the  **obsutil config -interactive**  command to set the access keys and endpoint of OBS.

    ```
    obsutil config -interactive
    
    Please input your ak:
    xxxxxxxxxxxxxxxxxxxxxxxxx
    Please input your sk:
    xxxxxxxxxxxxxxxxxxxxxxxxx
    Please input your endpoint:
    xxxxxxxxxxxxxxxxxxxxxxxxx
    Please input your token:
    xxxxxxxxxxxxxxxxxxxxxxxxx
    Config file url:
      C:\Users\tools\.obsutilconfig
    
    Update config file successfully!
    ```


