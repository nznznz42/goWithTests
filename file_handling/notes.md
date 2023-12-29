# chapter 17

- the testing/fstest package provides utilities to test filesystems
- the fs package defines basic interfaces to a file system
- the filesystem is usually provided by the OS but can also be provided by other software or packages
- some useful functions provided by fs are:
    - func ReadFile(fsys FS, name string)
    - func WalkDir(fsys FS, root string, fn WalkDirFunc) 
    - fs.open(name String)
