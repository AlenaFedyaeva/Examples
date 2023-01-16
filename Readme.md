# Examples 

## gRPC 
[gRPC example here](gRPC/README.md)

## Cobra 
[Cobra example here](cobra/cobra.md)

## REST

### 1) force download

```

func ForceDownload(w http.ResponseWriter, r *http.Request) {

         downloadBytes, err := ioutil.ReadFile(file)

         if err != nil {
                 fmt.Println(err)
         }

         // set the default MIME type to send
         mime := http.DetectContentType(downloadBytes)

         fileSize := len(string(downloadBytes))

         // Generate the server headers
         w.Header().Set("Content-Type", mime)
         w.Header().Set("Content-Disposition", "attachment; filename="+file+"")
         w.Header().Set("Expires", "0")
         w.Header().Set("Content-Transfer-Encoding", "binary")
         w.Header().Set("Content-Length", strconv.Itoa(fileSize))
         w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

         //b := bytes.NewBuffer(downloadBytes)
         //if _, err := b.WriteTo(w); err != nil {
         //              fmt.Fprintf(w, "%s", err)
         //      }

         // force it down the client's.....
         http.ServeContent(w, r, file, time.Now(), bytes.NewReader(downloadBytes))

 }
 
```
