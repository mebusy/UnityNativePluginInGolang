using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.Runtime.InteropServices;

public class GoPluginTest : MonoBehaviour
{
	[DllImport ("OSXBundle")]
	private static extern int lib_Add (int a,int b);
    // Start is called before the first frame update
    void Start()
    {
        Debug.Log( "test start" ) ;
        Debug.Log( lib_Add(1,2) ) ;
    }

    // Update is called once per frame
    void Update()
    {
        
    }
}
