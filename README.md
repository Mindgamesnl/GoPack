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
to 1.11 (remove secrets): 100.00% [================] 328 p/s 1m8s
INFO[0071] out/1.11-1.12.zip turned out to be 15MB and contains 2459 of 2472 
to 1.13 (remove secrets, flattening): 100.00% [================] 2911 p/s 1m19s
INFO[0152] out/1.13-1.14.zip turned out to be 15MB and contains 2472 of 4886 
to 1.15 (remove secrets, flattening): 100.00% [================] 44027 p/s 1m16s
INFO[0230] out/1.15.zip turned out to be 15MB and contains 2471 of 2568 
to 1.16 (remove secrets, flattening): 100.00% [================] 46270 p/s 1m12s
INFO[0304] out/1.16.zip turned out to be 15MB and contains 2471 of 2472 
INFO[0304] Finished pipeline, cleaning up..             
``` 

# TODO:
- Subedive particle textures