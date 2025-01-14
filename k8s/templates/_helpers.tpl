
{{- define "flightservice.fullname" -}}
{{- $name := default "flightservice" .Values.nameOverride -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "flightservice.name" -}}
{{- default "flightservice" .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

