package com.inkeliz.example.directory;

import java.lang.Runnable;
import java.lang.String;
import java.io.File;
import android.os.Environment;
import android.content.Context;
import android.content.ClipboardManager;
import android.content.ClipData;
import android.content.Intent;
import android.util.Log;
import android.app.Activity;
import android.app.Fragment;
import android.app.FragmentManager;
import android.app.FragmentTransaction;
import android.app.PendingIntent;
import android.content.IntentSender;
import android.content.IntentSender.OnFinished;
import android.content.IntentSender.SendIntentException;
import android.os.Bundle;
import android.Manifest;
import android.content.pm.PackageManager;

// That import is used to call GOJNI
import github.com.inkeliz.gojni.register_android;

public class directory_android extends Fragment {
	Context ctx;
	final int PERMISSIONS_REQUEST = 1;

	public directory_android() {
	}

	@Override public void onAttach(Context ctx) {
		super.onAttach(ctx);
		this.ctx = ctx;

        AskPermission();

        // That is used to call GOJNI and make directory_android accessible from Golang
        register_android.Register(this);
	}

	public String Pictures() {
	    AskPermission();
        return Environment.getExternalStoragePublicDirectory(Environment.DIRECTORY_DCIM).getAbsolutePath();
    }

	public String Download() {
	    AskPermission();
        return Environment.getExternalStoragePublicDirectory(Environment.DIRECTORY_DOWNLOADS).getAbsolutePath();
    }

	public String Music() {
	    AskPermission();
        return Environment.getExternalStoragePublicDirectory(Environment.DIRECTORY_MUSIC).getAbsolutePath();
    }

	public String Video() {
	    AskPermission();
        return Environment.getExternalStoragePublicDirectory(Environment.DIRECTORY_MOVIES).getAbsolutePath();
    }

	public String Document() {
	    AskPermission();
        return Environment.getExternalStoragePublicDirectory(Environment.DIRECTORY_DOCUMENTS).getAbsolutePath();
    }

    public void AskPermission() {
        if (ctx.checkSelfPermission(Manifest.permission.READ_EXTERNAL_STORAGE) != PackageManager.PERMISSION_GRANTED || ctx.checkSelfPermission(Manifest.permission.WRITE_EXTERNAL_STORAGE) != PackageManager.PERMISSION_GRANTED) {
                requestPermissions(new String[]{Manifest.permission.READ_EXTERNAL_STORAGE, Manifest.permission.WRITE_EXTERNAL_STORAGE}, PERMISSIONS_REQUEST);
        }
    }
}