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
INFO[0000] pack.zip currently is 18MB                   
INFO[0000] Loaded pack: Wakanda Forever. (1977-2020) in format 3 
to 1.11 (remove secrets): 100.00% [======================================================================================================================================================================================] 228 p/s 1m16s
INFO[0077] Flushing 2472 out of 2472 files              
INFO[0079] out/1.11-1.12.zip turned out to be 15MB      
to 1.13 (remove secrets, flattening): 100.00% [=========================================================================================================================================================================] 2916 p/s 1m16s
INFO[0155] Flushing 2472 out of 4881 files              
INFO[0158] out/1.13-1.14.zip turned out to be 15MB      
to 1.15 (remove secrets, flattening): 100.00% [========================================================================================================================================================================] 41052 p/s 1m20s
INFO[0239] Flushing 2471 out of 2567 files              
INFO[0241] out/1.15.zip turned out to be 15MB           
to 1.16 (remove secrets, flattening): 100.00% [========================================================================================================================================================================] 41568 p/s 1m20s
INFO[0322] Flushing 2471 out of 2472 files              
INFO[0324] out/1.16.zip turned out to be 15MB           
INFO[0324] Finished pipeline, cleaning up..  
``` 

# TODO:
- Subedive particle textures