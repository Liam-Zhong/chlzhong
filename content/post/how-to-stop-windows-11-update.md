+++
author = "Tuffy"
title = 'How to Stop Windows 11 Update'
date = 2024-12-15T20:59:13+08:00
math = true 
draft = false
comments = true
toc = false
description = "How to Stop Windows 11 Update"
+++

Hey, Are You Frustrated with Windows Automatic Updates? Here Are Some Methods to Stop Them :

Via Services Settings

- Open "Run"
- Type `services.msc`
- Find "Windows Update" service
- Set Startup Type to "Disabled"

Using Group Policy

- Press Win + R to open "Run"
- Type `gpedit.msc`
- Navigate to: Computer Configuration > Administrative Templates > Windows Components > Windows Update
- Configure "Configure Automatic Updates"

Temporarily Pause Updates

- Settings > Update & Security > Windows Update
- Choose "Pause Updates"
- Can pause up to 35 days

Using Registry

- Press Win + R to open "Run"
- Type `regedit`
- Navigate to `HKEY_LOCAL_MACHINE\SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate\AU`
- Create DWORD value `NoAutoUpdate`, set to 1

And you do like this,when you restart,you may say "what f**k man?",So Don't Do Like Above.

**Just press `田` , then input `services` , find the `Update Orchestrator Service` , change the status to `disabled` and set the `Recovery` always `Take No Action`**. :P
