{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "cert-manager.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "cert-manager.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "cert-manager.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "cert-manager.serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
    {{ default (include "cert-manager.fullname" .) .Values.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{/*
Create the default PodDisruptionBudget to use
*/}}
{{- define "podDisruptionBudget.spec" -}}
{{- if and .Values.podDisruptionBudget.minAvailable .Values.podDisruptionBudget.maxUnavailable }}
{{- fail "Cannot set both .Values.podDisruptionBudget.minAvailable and .Values.podDisruptionBudget.maxUnavailable" -}}
{{- end }}
{{- if not .Values.podDisruptionBudget.maxUnavailable }}
minAvailable: {{ default 1 .Values.podDisruptionBudget.minAvailable }}
{{- end }}
{{- if .Values.podDisruptionBudget.maxUnavailable }}
maxUnavailable: {{ .Values.podDisruptionBudget.maxUnavailable }}
{{- end }}
{{- end }}

{{/*
Webhook templates
*/}}

{{/*
Create the PodDisruptionBudget to use
*/}}
{{- define "webhook.podDisruptionBudget.spec" -}}
{{- if and .Values.webhook.podDisruptionBudget.minAvailable .Values.webhook.podDisruptionBudget.maxUnavailable }}
{{- fail "Cannot set both .Values.webhook.podDisruptionBudget.minAvailable and .Values.webhook.podDisruptionBudget.maxUnavailable" -}}
{{- end }}
{{- if not .Values.webhook.podDisruptionBudget.maxUnavailable }}
minAvailable: {{ default 1 .Values.webhook.podDisruptionBudget.minAvailable }}
{{- end }}
{{- if .Values.webhook.podDisruptionBudget.maxUnavailable }}
maxUnavailable: {{ .Values.webhook.podDisruptionBudget.maxUnavailable }}
{{- end }}
{{- end }}

{{/*
Expand the name of the chart.
Manually fix the 'app' and 'name' labels to 'webhook' to maintain
compatibility with the v0.9 deployment selector.
*/}}
{{- define "webhook.name" -}}
{{- printf "webhook" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "webhook.fullname" -}}
{{- $trimmedName := printf "%s" (include "cert-manager.fullname" .) | trunc 55 | trimSuffix "-" -}}
{{- printf "%s-webhook" $trimmedName | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "webhook.caRef" -}}
{{ .Release.Namespace}}/{{ template "webhook.fullname" . }}-ca
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "webhook.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "webhook.serviceAccountName" -}}
{{- if .Values.webhook.serviceAccount.create -}}
    {{ default (include "webhook.fullname" .) .Values.webhook.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.webhook.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{/*
cainjector templates
*/}}

{{/*
Create the PodDisruptionBudget to use
*/}}
{{- define "cainjector.podDisruptionBudget.spec" -}}
{{- if and .Values.cainjector.podDisruptionBudget.minAvailable .Values.cainjector.podDisruptionBudget.maxUnavailable }}
{{- fail "Cannot set both .Values.cainjector.podDisruptionBudget.minAvailable and .Values.cainjector.podDisruptionBudget.maxUnavailable" -}}
{{- end }}
{{- if not .Values.cainjector.podDisruptionBudget.maxUnavailable }}
minAvailable: {{ default 1 .Values.cainjector.podDisruptionBudget.minAvailable }}
{{- end }}
{{- if .Values.cainjector.podDisruptionBudget.maxUnavailable }}
maxUnavailable: {{ .Values.cainjector.podDisruptionBudget.maxUnavailable }}
{{- end }}
{{- end }}

{{/*
Expand the name of the chart.
Manually fix the 'app' and 'name' labels to 'cainjector' to maintain
compatibility with the v0.9 deployment selector.
*/}}
{{- define "cainjector.name" -}}
{{- printf "cainjector" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "cainjector.fullname" -}}
{{- $trimmedName := printf "%s" (include "cert-manager.fullname" .) | trunc 52 | trimSuffix "-" -}}
{{- printf "%s-cainjector" $trimmedName | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "cainjector.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "cainjector.serviceAccountName" -}}
{{- if .Values.cainjector.serviceAccount.create -}}
    {{ default (include "cainjector.fullname" .) .Values.cainjector.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.cainjector.serviceAccount.name }}
{{- end -}}
{{- end -}}
