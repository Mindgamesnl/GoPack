<p align="center">
    <img src="https://i.imgur.com/hXbjPIC.png" border="0">
</p>

##### Fast and native minecraft resourcepack translator in Go

GoPack automatically compiles `1.12`, `1.13`, `1.14`, `1.15` and `1.16` resourcepacks from a single `1.12.2` resource pack. It does this by completely scanning the input pack and then running migrations to rename all the assets and update the json pointers. It also removes any potential secrets during this process and can even minimize assets. GoPack supports Textures, Models, Sounds, Asset Compression and Language Files. 

# How to use
1. Download the latest [release](https://github.com/Mindgamesnl/GoPack/releases) for your platform.
2. Place the binary next to your 1.12 resourcepack, which must be named `pack.zip`
3. Execute with `./GoPack` or `./GoPack.exe`
4. That's it! give it some time and watch your packs appear in the `out/` directory.

### Example output:
```
mindgamesnl@redstaros-3 GoPack % go run main.go
INFO[0001] Loaded pack: Wakanda Forever. (1977-2020) in format 3 
INFO[0001] Starting pipelines                           
INFO[0001] Initializing pipelines took 0MS              
INFO[0001] Executing pipeline: to 1.11 (remove secrets) 
INFO[0001] Loading files into memory...                 
INFO[0001] Loaded 2446 files                            
12230 / 12230 [--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------] 100.00% 159 p/s 1m17s
INFO[0078] Flushing 2446 files                          
INFO[0080] Converting done. Validating written files... 
INFO[0080] Files seem OK                                
INFO[0080] Executing pipeline: to 1.13 (remove secrets, flattening) 
INFO[0080] Loading files into memory...                 
INFO[0081] Loaded 2446 files                            
212802 / 212802 [-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------] 100.00% 3050 p/s 1m10s
INFO[0151] Flushing 2605 files                          
INFO[0152] Converting done. Validating written files... 
INFO[0153] Files seem OK                                
INFO[0153] Executing pipeline: to 1.15 (remove secrets, flattening) 
INFO[0153] Loading files into memory...                 
INFO[0154] Loaded 2446 files                            
3270302 / 3270302 [--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------] 100.00% 41862 p/s 1m18s
INFO[0232] Flushing 2446 files                          
INFO[0234] Converting done. Validating written files... 
INFO[0234] Files seem OK                                
INFO[0234] Executing pipeline: to 1.16 (remove secrets, flattening) 
INFO[0234] Loading files into memory...                 
INFO[0235] Loaded 2446 files                            
3270302 / 3270302 [--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------] 100.00% 44135 p/s 1m15s
INFO[0309] Flushing 2446 files                          
INFO[0311] Converting done. Validating written files... 
INFO[0311] Files seem OK
``` 

# TODO:
- Subedive particle textures

### Credits
- Png compression by [lossy](https://github.com/foobaz/lossypng) (slightly modified to support in-memory compression)