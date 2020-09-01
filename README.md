# GoPack
Minecraft resourcepack translator in Go

GoPack automatically compiles `1.12`, `1.13`, `1.14`, `1.15` and `1.16` resourcepacks from a single `1.12.2` resource pack. It does this by completely scanning the input pack and then running migrations to rename all the assets and update the json pointers. It also removes any potential secrets during this process and can even minimize assets. GoPack supports Textures, Models and Sounds.

Every target pack takes around 1 minute to compile and generally executes `2502258` tasks *(on the ImagineFun pack)*. Multi threading is work in progress, it won't really shorten the individual time per pack but should allow you to compile multiple packs at once.

### Example output:
```
INFO[0000] Loaded pack: Wakanda Forever. (1977-2020) in format 3
INFO[0000] Starting pipelines
INFO[0000] Initializing pipelines took 0MS
INFO[0000] Executing pipeline: to 1.11 (remove secrets)
12230 / 12230 [---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------] 100.00% 191 p/s 1m4s
INFO[0065] Flushing files..
INFO[0065] Saved!
INFO[0065] Executing pipeline: to 1.13 (remove secrets, flattening)
2502258 / 2502258 [---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------] 100.00% 39222 p/s 1m4s
INFO[0129] Flushing files..
INFO[0130] Saved!
INFO[0130] Executing pipeline: to 1.15 (remove secrets, flattening)
2502258 / 2502258 [---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------] 100.00% 39990 p/s 1m3s
INFO[0193] Flushing files..
INFO[0193] Saved!
INFO[0193] Executing pipeline: to 1.16 (remove secrets, flattening)
2502258 / 2502258 [---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------] 100.00% 41519 p/s 1m0s
INFO[0254] Flushing files..
INFO[0254] Saved!
``` 