package source

import (
	"github.com/openshift/cluster-logging-operator/internal/generator"
)

const JournalLogTemplate = `
{{define "inputSourceJournalTemplate" -}}
# {{.Desc}}
<source>
  @type systemd
  @id systemd-input
  @label @{{.OutLabel}}
  path '/var/log/journal'
  <storage>
    @type local
    persistent true
    # NOTE: if this does not end in .json, fluentd will think it
    # is the name of a directory - see fluentd storage_local.rb
    path '/var/lib/fluentd/pos/journal_pos.json'
  </storage>
  matches "#{ENV['JOURNAL_FILTERS_JSON'] || '[]'}"
  tag journal
  read_from_head "#{if (val = ENV.fetch('JOURNAL_READ_FROM_HEAD','')) && (val.length > 0); val; else 'false'; end}"
</source>
{{end}}`

type JournalLog = generator.ConfLiteral
