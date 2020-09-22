using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.Runtime.InteropServices;

public class GoPluginTest : MonoBehaviour
{
	[DllImport ("OSXBundle")]
	private static extern int lib_Add (int a,int b);
    [DllImport ("OSXBundle")]
    private static extern int lib_SendMsg(string p0);


    [DllImport ("OSXBundle")]
    private static extern void lib_SetCallbackFunc(SimpleArgCallback p0);

    private delegate string SimpleArgCallback(string arg);
    private string SimpleArgCallbackUnityFun(string arg) {
        Debug.Log( "received callback with arg:" + arg );
        return "unity echo back:" + arg ;
    }



    // Start is called before the first frame update
    void Start()
    {
        Debug.Log( "test start" ) ;

        SimpleArgCallback sc = new SimpleArgCallback(SimpleArgCallbackUnityFun);
        lib_SetCallbackFunc(sc);

        Debug.Log( lib_Add(1,2) ) ;

        int ret = lib_SendMsg( "send from unity" );
        Debug.Log( "ret of SendMsg " + ret  ) ;
    }

    // Update is called once per frame
    void Update()
    {
        
    }


}
