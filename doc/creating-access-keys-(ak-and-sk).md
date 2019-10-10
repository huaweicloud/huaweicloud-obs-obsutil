# Creating Access Keys \(AK and SK\)<a name="EN-US_TOPIC_0172237146"></a>

This section describes how to create access keys \(AKs and SKs\) in OBS Console. A pair of AK and SK is used to encrypt the signature of a request, ensuring that the request is secure and integral, and that identities of the request sender and receiver are correct.

## Scenarios<a name="s6eee9c5cf28244198d6c28ef50ce2276"></a>

AKs and SKs support the authentication mechanism of Identity and Access Management \(IAM\).

-   Access key ID \(AK\): indicates the ID of the access key, which is a unique identifier used together with a secret access key to sign requests cryptographically.
-   Secret access key \(SK\): indicates the private key used together with its associated AK to cryptographically sign requests. The AK and SK are used together to identify a request sender to prevent the request from being modified.

## Restrictions and Limitations<a name="section64691490143136"></a>

Each user can create up to two valid AK/SK pairs.

## Prerequisites<a name="section37631452155356"></a>

An account has been registered and activated.

## Procedure<a name="section13220848205019"></a>

1.  In the upper right corner of the console page, select  **My Credential**  under the username.
2.  On the  **My Credentials**  page, choose  **Access Keys**  \>  **Add Access Key**.
3.  In the  **Add Access Key**  dialog box that is displayed, enter the password and its verification code.

    >![](public_sys-resources/icon-note.gif) **NOTE:**   
    >-   If you have not bound an email address or mobile number, enter only the password.  
    >-   If you have bound an email address and a mobile number, you can select the verification by email or mobile phone.  

4.  Click  **OK**.
5.  In the  **Download Access Key**  dialog box that is displayed, click  **OK**  to save the access keys to your browser's default download path.

    >![](public_sys-resources/icon-note.gif) **NOTE:**   
    >Keep the access keys properly to prevent information leakage. If you click  **Cancel**  in the dialog box, the access keys will not be downloaded, and you cannot download then later. Re-create access keys if required.  

6.  Open the downloaded  **credentials.csv**  file to obtain the access keys \(AK and SK\).

    >![](public_sys-resources/icon-note.gif) **NOTE:**   
    >In the access key file, the value in the  **Access Key ID**  column is the AK, and the value in the  **Secret Access Key**  column is the SK.  


