package indexmanagement

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/ViaQ/logerr/log"
	logging "github.com/openshift/cluster-logging-operator/apis/logging/v1"
	esapi "github.com/openshift/elasticsearch-operator/apis/logging/v1"
)

const (
	PolicyNameApp   = "app-policy"
	PolicyNameInfra = "infra-policy"
	PolicyNameAudit = "audit-policy"

	MappingNameApp   = "app"
	MappingNameInfra = "infra"
	MappingNameAudit = "audit"

	PollInterval = "15m"

	HotPhaseAgeAsPercentOfMaxAge = 5
)

var (
	AliasesApp   = []string{"app", "logs.app"}
	AliasesInfra = []string{"infra", "logs.infra"}
	AliasesAudit = []string{"audit", "logs.audit"}

	agePattern = regexp.MustCompile(`^(?P<number>\d+)(?P<unit>[yMwdhHms])$`)
)

func NewSpec(retentionPolicy *logging.RetentionPoliciesSpec) *esapi.IndexManagementSpec {

	retentionPolicy = newDefaultPoliciesSpec(retentionPolicy)
	indexManagement := esapi.IndexManagementSpec{}
	if retentionPolicy.App != nil {
		hotPhaseAgeApp, err := getHotPhaseAge(retentionPolicy.App.MaxAge)
		if err != nil {
			log.Error(err, "Error occurred while getting hot phase age for App log source")
			return nil
		}
		appPolicySpec := newPolicySpec(PolicyNameApp, retentionPolicy.App, hotPhaseAgeApp)
		indexManagement.Policies = append(indexManagement.Policies, appPolicySpec)
		appMappingSpec := newMappingSpec(MappingNameApp, PolicyNameApp, AliasesApp)
		indexManagement.Mappings = append(indexManagement.Mappings, appMappingSpec)
	}
	if retentionPolicy.Infra != nil {
		hotPhaseAgeInfra, err := getHotPhaseAge(retentionPolicy.Infra.MaxAge)
		if err != nil {
			log.Error(err, "Error occurred while getting hot phase age for Infra log source.")
			return nil
		}
		infraPolicySpec := newPolicySpec(PolicyNameInfra, retentionPolicy.Infra, hotPhaseAgeInfra)
		indexManagement.Policies = append(indexManagement.Policies, infraPolicySpec)
		infraMappingSpec := newMappingSpec(MappingNameInfra, PolicyNameInfra, AliasesInfra)
		indexManagement.Mappings = append(indexManagement.Mappings, infraMappingSpec)
	}
	if retentionPolicy.Audit != nil {
		hotPhaseAgeAudit, err := getHotPhaseAge(retentionPolicy.Audit.MaxAge)
		if err != nil {
			log.Error(err, "Error occurred while getting hot phase age for Audit log source.")
			return nil
		}
		auditPolicySpec := newPolicySpec(PolicyNameAudit, retentionPolicy.Audit, hotPhaseAgeAudit)
		indexManagement.Policies = append(indexManagement.Policies, auditPolicySpec)
		auditMappingSpec := newMappingSpec(MappingNameAudit, PolicyNameAudit, AliasesAudit)
		indexManagement.Mappings = append(indexManagement.Mappings, auditMappingSpec)
	}
	return &indexManagement
}

func newDefaultPoliciesSpec(spec *logging.RetentionPoliciesSpec) *logging.RetentionPoliciesSpec {

	defaultSpec := &logging.RetentionPoliciesSpec{
		App: &logging.RetentionPolicySpec{
			MaxAge:                  esapi.TimeUnit("7d"),
			PruneNamespacesInterval: esapi.TimeUnit("30m"),
		},
		Infra: &logging.RetentionPolicySpec{
			MaxAge:                  esapi.TimeUnit("7d"),
			PruneNamespacesInterval: esapi.TimeUnit("30m"),
		},
		Audit: &logging.RetentionPolicySpec{
			MaxAge:                  esapi.TimeUnit("7d"),
			PruneNamespacesInterval: esapi.TimeUnit("30m"),
		},
	}

	if spec != nil {
		if spec.App != nil {
			if spec.App.MaxAge != "" {
				defaultSpec.App.MaxAge = spec.App.MaxAge
			}
			if spec.App.Namespaces != nil {
				if spec.App.PruneNamespacesInterval != "" {
					defaultSpec.App.PruneNamespacesInterval = spec.App.PruneNamespacesInterval
				}
				defaultSpec.App.Namespaces = spec.App.Namespaces
			}
		}
		if spec.Infra != nil {
			if spec.Infra.MaxAge != "" {
				defaultSpec.Infra.MaxAge = spec.Infra.MaxAge
			}
			if spec.Infra.Namespaces != nil {
				if spec.Infra.PruneNamespacesInterval != "" {
					defaultSpec.Infra.PruneNamespacesInterval = spec.Infra.PruneNamespacesInterval
				}
				defaultSpec.Infra.Namespaces = spec.Infra.Namespaces
			}
		}
		if spec.Audit != nil {
			if spec.Audit.MaxAge != "" {
				defaultSpec.Audit.MaxAge = spec.Audit.MaxAge
			}
			if spec.Audit.Namespaces != nil {
				if spec.Audit.PruneNamespacesInterval != "" {
					defaultSpec.Audit.PruneNamespacesInterval = spec.Audit.PruneNamespacesInterval
				}
				defaultSpec.Audit.Namespaces = spec.Audit.Namespaces
			}
		}
	}

	return defaultSpec
}

func newPolicySpec(name string, retentionPolicy *logging.RetentionPolicySpec, hotPhaseAge esapi.TimeUnit) esapi.IndexManagementPolicySpec {

	policySpec := esapi.IndexManagementPolicySpec{
		Name:         name,
		PollInterval: PollInterval,
		Phases: esapi.IndexManagementPhasesSpec{
			Hot: &esapi.IndexManagementHotPhaseSpec{
				Actions: esapi.IndexManagementActionsSpec{
					Rollover: &esapi.IndexManagementActionSpec{
						MaxAge: hotPhaseAge,
					},
				},
			},
			Delete: &esapi.IndexManagementDeletePhaseSpec{
				MinAge:                  retentionPolicy.MaxAge,
				PruneNamespacesInterval: retentionPolicy.PruneNamespacesInterval,
			},
		},
	}

	if retentionPolicy.Namespaces != nil {
		policySpec.Phases.Delete.Namespaces = retentionPolicy.Namespaces
	}

	return policySpec
}

func newMappingSpec(name string, policyRef string, aliases []string) esapi.IndexManagementPolicyMappingSpec {
	mappingSpec := esapi.IndexManagementPolicyMappingSpec{
		Name:      name,
		PolicyRef: policyRef,
		Aliases:   aliases,
	}
	return mappingSpec
}

func getHotPhaseAge(maxAge esapi.TimeUnit) (esapi.TimeUnit, error) {
	var (
		age         int
		unit        byte
		err         error
		hotphaseAge int
	)
	age, unit, err = toAgeAndUnit(maxAge)
	if err != nil {
		return esapi.TimeUnit(""), err
	}

	if age == 0 {
		return esapi.TimeUnit(fmt.Sprintf("0%c", unit)), nil
	}

	hotphaseAge, unit, err = toHotPhaseAge(age, unit)
	if err != nil {
		return esapi.TimeUnit(""), err
	}

	return esapi.TimeUnit(fmt.Sprintf("%d%c", hotphaseAge, unit)), nil
}

func toAgeAndUnit(timeunit esapi.TimeUnit) (int, byte, error) {
	strvalues := agePattern.FindStringSubmatch(string(timeunit))
	if len(strvalues) != 3 {
		return 0, 0, fmt.Errorf("age pattern mismatch")
	}
	age, _ := strconv.Atoi(strvalues[1])
	unit := strvalues[2][0]
	return age, unit, nil
}

func toHotPhaseAge(value int, unit byte) (int, byte, error) {
	newval := value * HotPhaseAgeAsPercentOfMaxAge / 100

	for newval == 0 {
		value, newunit, err := convertToLowerUnits(value, unit)
		if err != nil {
			return 0, 0, err
		}
		newval = value * HotPhaseAgeAsPercentOfMaxAge / 100
		unit = newunit
	}

	return newval, unit, nil
}

func convertToLowerUnits(value int, unit byte) (int, byte, error) {

	switch unit {
	case 's':
		return 0, 0, fmt.Errorf("cannot convert \"%d%c\" to lower units", value, unit)
	case 'm':
		newval := value * 60
		return newval, 's', nil
	case 'h', 'H':
		newval := value * 60
		return newval, 'm', nil
	case 'd':
		newval := value * 24
		return newval, 'h', nil
	case 'w':
		newval := value * 7
		return newval, 'd', nil
	case 'M':
		newval := value * 30
		return newval, 'd', nil
	case 'y':
		newval := value * 365
		return newval, 'd', nil
	}

	return 0, 0, fmt.Errorf("unknown units")
}
