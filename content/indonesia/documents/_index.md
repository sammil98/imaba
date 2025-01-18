---
title: "Dokumen Organisasi"
description: "Berikut adalah daftar dokumen yang dapat Anda akses langsung di halaman ini."
---

{{< tabs >}}

{{ $pdfPath := "assets/document" }}
{{ range readDir $pdfPath }}
{{< tab (replace .Name ".pdf" "") >}}
#### {{ .Name }}
<iframe src="/pdf/{{ .Name }}" width="100%" height="600px" style="border: none;"></iframe>
{{< /tab >}}
{{ end }}

{{< /tabs >}}
