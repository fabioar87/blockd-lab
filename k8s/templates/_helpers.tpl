
{{- define "transactionservice.fullname" -}}
{{- $name := default "transactionservice" .Values.nameOverride -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "transactionservice.name" -}}
{{- default "transactionservice" .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}
