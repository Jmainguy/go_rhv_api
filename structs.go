package main

type RhevAuth struct {
	AccessToken string `json:"access_token"`
	Exp         string `json:"exp"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type StorageDomains struct {
	StorageDomain []struct {
		Available                  int64  `json:"available,omitempty"`
		Committed                  int    `json:"committed"`
		CriticalSpaceActionBlocker string `json:"critical_space_action_blocker"`
		DiscardAfterDelete         string `json:"discard_after_delete"`
		ExternalStatus             string `json:"external_status"`
		Master                     string `json:"master"`
		Storage                    struct {
			Address string `json:"address"`
			Path    string `json:"path"`
			Type    string `json:"type"`
		} `json:"storage"`
		StorageFormat             string `json:"storage_format"`
		SupportsDiscard           string `json:"supports_discard"`
		SupportsDiscardZeroesData string `json:"supports_discard_zeroes_data"`
		Type                      string `json:"type"`
		Used                      int64  `json:"used,omitempty"`
		WarningLowSpaceIndicator  string `json:"warning_low_space_indicator"`
		WipeAfterDelete           string `json:"wipe_after_delete"`
		DataCenters               struct {
			DataCenter []struct {
				ID string `json:"id"`
			} `json:"data_center"`
		} `json:"data_centers,omitempty"`
		Actions struct {
			Link []struct {
				Href string `json:"href"`
				Rel  string `json:"rel"`
			} `json:"link"`
		} `json:"actions"`
		Name string `json:"name"`
		Href string `json:"href"`
		ID   string `json:"id"`
		Link []struct {
			Href string `json:"href"`
			Rel  string `json:"rel"`
		} `json:"link"`
		Status      string `json:"status,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"storage_domain"`
}

type RhevHosts struct {
	Host []struct {
		Address        string `json:"address"`
		AutoNumaStatus string `json:"auto_numa_status"`
		Certificate    struct {
			Organization string `json:"organization"`
			Subject      string `json:"subject"`
		} `json:"certificate"`
		CPU struct {
			Name     string `json:"name"`
			Speed    int    `json:"speed"`
			Topology struct {
				Cores   string `json:"cores"`
				Sockets string `json:"sockets"`
				Threads string `json:"threads"`
			} `json:"topology"`
		} `json:"cpu"`
		DevicePassthrough struct {
			Enabled string `json:"enabled"`
		} `json:"device_passthrough"`
		ExternalStatus      string `json:"external_status"`
		HardwareInformation struct {
			Manufacturer        string `json:"manufacturer"`
			ProductName         string `json:"product_name"`
			SerialNumber        string `json:"serial_number"`
			SupportedRngSources struct {
				SupportedRngSource []string `json:"supported_rng_source"`
			} `json:"supported_rng_sources"`
			UUID string `json:"uuid"`
		} `json:"hardware_information"`
		Iscsi struct {
			Initiator string `json:"initiator"`
		} `json:"iscsi"`
		KdumpStatus string `json:"kdump_status"`
		Ksm         struct {
			Enabled string `json:"enabled"`
		} `json:"ksm"`
		LibvirtVersion struct {
			Build       string `json:"build"`
			FullVersion string `json:"full_version"`
			Major       string `json:"major"`
			Minor       string `json:"minor"`
			Revision    string `json:"revision"`
		} `json:"libvirt_version"`
		MaxSchedulingMemory int64  `json:"max_scheduling_memory"`
		Memory              int64  `json:"memory"`
		NumaSupported       string `json:"numa_supported"`
		Os                  struct {
			CustomKernelCmdline   string `json:"custom_kernel_cmdline"`
			ReportedKernelCmdline string `json:"reported_kernel_cmdline"`
			Type                  string `json:"type"`
			Version               struct {
				FullVersion string `json:"full_version"`
				Major       string `json:"major"`
				Minor       string `json:"minor"`
			} `json:"version"`
		} `json:"os"`
		Port            string `json:"port"`
		PowerManagement struct {
			AutomaticPmEnabled string `json:"automatic_pm_enabled"`
			Enabled            string `json:"enabled"`
			KdumpDetection     string `json:"kdump_detection"`
			PmProxies          struct {
				PmProxy []struct {
					Type string `json:"type"`
				} `json:"pm_proxy"`
			} `json:"pm_proxies"`
		} `json:"power_management"`
		Protocol string `json:"protocol"`
		SeLinux  struct {
			Mode string `json:"mode"`
		} `json:"se_linux"`
		Spm struct {
			Priority string `json:"priority"`
			Status   string `json:"status"`
		} `json:"spm"`
		SSH struct {
			Fingerprint string `json:"fingerprint"`
			Port        string `json:"port"`
		} `json:"ssh"`
		Status  string `json:"status"`
		Summary struct {
			Active    string `json:"active"`
			Migrating string `json:"migrating"`
			Total     string `json:"total"`
		} `json:"summary"`
		TransparentHugepages struct {
			Enabled string `json:"enabled"`
		} `json:"transparent_hugepages"`
		Type            string `json:"type"`
		UpdateAvailable string `json:"update_available"`
		Version         struct {
			Build       string `json:"build"`
			FullVersion string `json:"full_version"`
			Major       string `json:"major"`
			Minor       string `json:"minor"`
			Revision    string `json:"revision"`
		} `json:"version"`
		Cluster struct {
			Href string `json:"href"`
			ID   string `json:"id"`
		} `json:"cluster"`
		Actions struct {
			Link []struct {
				Href string `json:"href"`
				Rel  string `json:"rel"`
			} `json:"link"`
		} `json:"actions"`
		Name    string `json:"name"`
		Comment string `json:"comment"`
		Href    string `json:"href"`
		ID      string `json:"id"`
		Link    []struct {
			Href string `json:"href"`
			Rel  string `json:"rel"`
		} `json:"link"`
	} `json:"host"`
}

type Clusters struct {
	Cluster []struct {
		BallooningEnabled string `json:"ballooning_enabled"`
		CPU               struct {
			Architecture string `json:"architecture"`
			Type         string `json:"type"`
		} `json:"cpu"`
		CustomSchedulingPolicyProperties struct {
			Property []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"property"`
		} `json:"custom_scheduling_policy_properties"`
		ErrorHandling struct {
			OnError string `json:"on_error"`
		} `json:"error_handling"`
		FencingPolicy struct {
			Enabled                  string `json:"enabled"`
			SkipIfConnectivityBroken struct {
				Enabled   string `json:"enabled"`
				Threshold string `json:"threshold"`
			} `json:"skip_if_connectivity_broken"`
			SkipIfGlusterBricksUp     string `json:"skip_if_gluster_bricks_up"`
			SkipIfGlusterQuorumNotMet string `json:"skip_if_gluster_quorum_not_met"`
			SkipIfSdActive            struct {
				Enabled string `json:"enabled"`
			} `json:"skip_if_sd_active"`
		} `json:"fencing_policy"`
		GlusterService string `json:"gluster_service"`
		HaReservation  string `json:"ha_reservation"`
		Ksm            struct {
			Enabled          string `json:"enabled"`
			MergeAcrossNodes string `json:"merge_across_nodes"`
		} `json:"ksm"`
		MaintenanceReasonRequired string `json:"maintenance_reason_required"`
		MemoryPolicy              struct {
			OverCommit struct {
				Percent string `json:"percent"`
			} `json:"over_commit"`
			TransparentHugepages struct {
				Enabled string `json:"enabled"`
			} `json:"transparent_hugepages"`
		} `json:"memory_policy"`
		Migration struct {
			AutoConverge string `json:"auto_converge"`
			Bandwidth    struct {
				AssignmentMethod string `json:"assignment_method"`
			} `json:"bandwidth"`
			Compressed string `json:"compressed"`
			Policy     struct {
				ID string `json:"id"`
			} `json:"policy"`
		} `json:"migration"`
		OptionalReason     string `json:"optional_reason"`
		RequiredRngSources struct {
			RequiredRngSource []string `json:"required_rng_source"`
		} `json:"required_rng_sources"`
		SwitchType      string `json:"switch_type"`
		ThreadsAsCores  string `json:"threads_as_cores"`
		TrustedService  string `json:"trusted_service"`
		TunnelMigration string `json:"tunnel_migration"`
		Version         struct {
			Major string `json:"major"`
			Minor string `json:"minor"`
		} `json:"version"`
		VirtService string `json:"virt_service"`
		DataCenter  struct {
			Href string `json:"href"`
			ID   string `json:"id"`
		} `json:"data_center"`
		MacPool struct {
			Href string `json:"href"`
			ID   string `json:"id"`
		} `json:"mac_pool"`
		SchedulingPolicy struct {
			Href string `json:"href"`
			ID   string `json:"id"`
		} `json:"scheduling_policy"`
		Actions struct {
			Link []struct {
				Href string `json:"href"`
				Rel  string `json:"rel"`
			} `json:"link"`
		} `json:"actions"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Comment     string `json:"comment"`
		Href        string `json:"href"`
		ID          string `json:"id"`
		Link        []struct {
			Href string `json:"href"`
			Rel  string `json:"rel"`
		} `json:"link"`
	} `json:"cluster"`
}
