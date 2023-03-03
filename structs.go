package main

import "time"

type auth struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Expires      int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type me struct {
	Actions []struct {
		Href   string `json:"href"`
		Method string `json:"method"`
		Name   string `json:"name"`
	} `json:"actions"`
	BusinessUnit struct {
		_Type     string `json:"$type"`
		Addresses struct {
			_Type string        `json:"$type"`
			Href  string        `json:"href"`
			Items []interface{} `json:"items"`
			Total int           `json:"total"`
		} `json:"addresses"`
		BillingStorageType   int `json:"billingStorageType"`
		BillingStorageTypeID int `json:"billingStorageTypeId"`
		BrandConfig          struct {
			_Type       string    `json:"$type"`
			CreatedDate time.Time `json:"createdDate"`
			Href        string    `json:"href"`
			ID          int       `json:"id"`
			Logo        string    `json:"logo"`
			Styles      []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"styles"`
		} `json:"brandConfig"`
		BusinessUnits struct {
			_Type string        `json:"$type"`
			Href  string        `json:"href"`
			Items []interface{} `json:"items"`
			Total int           `json:"total"`
		} `json:"businessUnits"`
		ClientOptionSetFilter            []interface{} `json:"clientOptionSetFilter"`
		Consumers                        []interface{} `json:"consumers"`
		CreatedDate                      time.Time     `json:"createdDate"`
		DomainFilter                     []interface{} `json:"domainFilter"`
		FinalDeleteRequestApprover       bool          `json:"finalDeleteRequestApprover"`
		Href                             string        `json:"href"`
		ID                               int           `json:"id"`
		InvoiceDay                       int           `json:"invoiceDay"`
		Name                             string        `json:"name"`
		NodeLimit                        int           `json:"nodeLimit"`
		PasswordExpirationDays           int           `json:"passwordExpirationDays"`
		RegistrationNumber               string        `json:"registrationNumber"`
		ReportRemotely                   bool          `json:"reportRemotely"`
		RequireTotp                      bool          `json:"requireTotp"`
		RequiredApproversOfDeleteRequest int           `json:"requiredApproversOfDeleteRequest"`
		StorageLimit                     int           `json:"storageLimit"`
		SupportResponsible               bool          `json:"supportResponsible"`
		SystemID                         string        `json:"systemId"`
		Tags                             []interface{} `json:"tags"`
		TimeZone                         struct {
			_Type   string `json:"$type"`
			Actions []struct {
				Href   string `json:"href"`
				Method string `json:"method"`
				Name   string `json:"name"`
			} `json:"actions"`
			CreatedDate time.Time `json:"createdDate"`
			Href        string    `json:"href"`
			ID          int       `json:"id"`
			Name        string    `json:"name"`
			Offset      int       `json:"offset"`
			WindowsID   string    `json:"windowsId"`
		} `json:"timeZone"`
		TransferLimit       int  `json:"transferLimit"`
		UseScheduleBindings bool `json:"useScheduleBindings"`
		Users               struct {
			_Type string        `json:"$type"`
			Href  string        `json:"href"`
			Items []interface{} `json:"items"`
			Total int           `json:"total"`
		} `json:"users"`
		UsersCanApproveOwnRequests bool `json:"usersCanApproveOwnRequests"`
	} `json:"businessUnit"`
	BusinessUnitID                    int       `json:"businessUnitId"`
	CreatedDate                       time.Time `json:"createdDate"`
	CsvDelimiter                      string    `json:"csvDelimiter"`
	DateAtTimeFormat                  string    `json:"dateAtTimeFormat"`
	DateFormatInitialized             string    `json:"dateFormatInitialized"`
	DateTimeFormat                    string    `json:"dateTimeFormat"`
	DecimalMark                       string    `json:"decimalMark"`
	DisplayName                       string    `json:"displayName"`
	Email                             string    `json:"email"`
	ExcludeJobTags                    bool      `json:"excludeJobTags"`
	FailedLoginAttempts               int       `json:"failedLoginAttempts"`
	GroupStatusReportByBusinessUnits  bool      `json:"groupStatusReportByBusinessUnits"`
	Href                              string    `json:"href"`
	ID                                int       `json:"id"`
	ItemsPerPage                      int       `json:"itemsPerPage"`
	LastLoginDate                     time.Time `json:"lastLoginDate"`
	LastPasswordChange                time.Time `json:"lastPasswordChange"`
	Locked                            bool      `json:"locked"`
	Name                              string    `json:"name"`
	ReceiveCommentNotifications       bool      `json:"receiveCommentNotifications"`
	ReceiveDeleteRequestNotifications bool      `json:"receiveDeleteRequestNotifications"`
	Role                              struct {
		CreatedDate time.Time `json:"createdDate"`
		Description string    `json:"description"`
		ID          int       `json:"id"`
		Inheritable bool      `json:"inheritable"`
		Name        string    `json:"name"`
		Privileges  []struct {
			CreatedDate time.Time `json:"createdDate"`
			ID          int       `json:"id"`
			Name        string    `json:"name"`
			Operation   string    `json:"operation"`
			Resource    string    `json:"resource"`
		} `json:"privileges"`
	} `json:"role"`
	TimeFormatInitialized   string        `json:"timeFormatInitialized"`
	UseOwnReportErrorConfig bool          `json:"useOwnReportErrorConfig"`
	UserIdentities          []interface{} `json:"userIdentities"`
	UsesTotp                bool          `json:"usesTotp"`
}

type consumer struct {
	AllowNoActivity bool `json:"allowNoActivity"`
	BusinessUnit    struct {
		_Type     string `json:"$type"`
		Addresses struct {
			_Type string        `json:"$type"`
			Href  string        `json:"href"`
			Items []interface{} `json:"items"`
			Total int           `json:"total"`
		} `json:"addresses"`
		BillingStorageType   int `json:"billingStorageType"`
		BillingStorageTypeID int `json:"billingStorageTypeId"`
		BusinessUnits        struct {
			_Type string        `json:"$type"`
			Href  string        `json:"href"`
			Items []interface{} `json:"items"`
			Total int           `json:"total"`
		} `json:"businessUnits"`
		ClientOptionSetFilter            []interface{} `json:"clientOptionSetFilter"`
		Consumers                        []interface{} `json:"consumers"`
		CreatedDate                      time.Time     `json:"createdDate"`
		DomainFilter                     []interface{} `json:"domainFilter"`
		FinalDeleteRequestApprover       bool          `json:"finalDeleteRequestApprover"`
		Href                             string        `json:"href"`
		ID                               int           `json:"id"`
		InvoiceDay                       int           `json:"invoiceDay"`
		Name                             string        `json:"name"`
		NodeLimit                        int           `json:"nodeLimit"`
		PasswordExpirationDays           int           `json:"passwordExpirationDays"`
		RegistrationNumber               string        `json:"registrationNumber"`
		ReportRemotely                   bool          `json:"reportRemotely"`
		RequiredApproversOfDeleteRequest int           `json:"requiredApproversOfDeleteRequest"`
		StorageLimit                     int           `json:"storageLimit"`
		SupportResponsible               bool          `json:"supportResponsible"`
		Tags                             []interface{} `json:"tags"`
		TransferLimit                    int           `json:"transferLimit"`
		UseScheduleBindings              bool          `json:"useScheduleBindings"`
		Users                            struct {
			_Type string        `json:"$type"`
			Href  string        `json:"href"`
			Items []interface{} `json:"items"`
			Total int           `json:"total"`
		} `json:"users"`
		UsersCanApproveOwnRequests bool `json:"usersCanApproveOwnRequests"`
	} `json:"businessUnit"`
	CreatedDate                 time.Time `json:"createdDate"`
	DataSourceIsPotentialParent bool      `json:"dataSourceIsPotentialParent"`
	DataSourceState             struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"dataSourceState"`
	DataSourceType struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"dataSourceType"`
	Href string        `json:"href"`
	ID   int           `json:"id"`
	Name string        `json:"name"`
	Tags []interface{} `json:"tags"`
}

type node struct {
	ActivatedDate      time.Time     `json:"activatedDate"`
	ActivityLogFilters []interface{} `json:"activityLogFilters"`
	ActivityLogs       []interface{} `json:"activityLogs"`
	Consumer           struct {
		_Type           string `json:"$type"`
		AllowNoActivity bool   `json:"allowNoActivity"`
		BusinessUnit    struct {
			_Type     string `json:"$type"`
			Addresses struct {
				_Type string        `json:"$type"`
				Href  string        `json:"href"`
				Items []interface{} `json:"items"`
				Total int           `json:"total"`
			} `json:"addresses"`
			BillingStorageType   int `json:"billingStorageType"`
			BillingStorageTypeID int `json:"billingStorageTypeId"`
			BusinessUnits        struct {
				_Type string        `json:"$type"`
				Href  string        `json:"href"`
				Items []interface{} `json:"items"`
				Total int           `json:"total"`
			} `json:"businessUnits"`
			ClientOptionSetFilter            []interface{} `json:"clientOptionSetFilter"`
			Consumers                        []interface{} `json:"consumers"`
			CreatedDate                      time.Time     `json:"createdDate"`
			DomainFilter                     []interface{} `json:"domainFilter"`
			FinalDeleteRequestApprover       bool          `json:"finalDeleteRequestApprover"`
			Href                             string        `json:"href"`
			ID                               int           `json:"id"`
			InvoiceDay                       int           `json:"invoiceDay"`
			Name                             string        `json:"name"`
			NodeLimit                        int           `json:"nodeLimit"`
			PasswordExpirationDays           int           `json:"passwordExpirationDays"`
			RegistrationNumber               string        `json:"registrationNumber"`
			ReportRemotely                   bool          `json:"reportRemotely"`
			RequiredApproversOfDeleteRequest int           `json:"requiredApproversOfDeleteRequest"`
			StorageLimit                     int           `json:"storageLimit"`
			SupportResponsible               bool          `json:"supportResponsible"`
			Tags                             []interface{} `json:"tags"`
			TransferLimit                    int           `json:"transferLimit"`
			UseScheduleBindings              bool          `json:"useScheduleBindings"`
			Users                            struct {
				_Type string        `json:"$type"`
				Href  string        `json:"href"`
				Items []interface{} `json:"items"`
				Total int           `json:"total"`
			} `json:"users"`
			UsersCanApproveOwnRequests bool `json:"usersCanApproveOwnRequests"`
		} `json:"businessUnit"`
		CreatedDate                 time.Time `json:"createdDate"`
		DataSourceIsPotentialParent bool      `json:"dataSourceIsPotentialParent"`
		DataSourceState             struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"dataSourceState"`
		DataSourceType struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"dataSourceType"`
		Href string        `json:"href"`
		ID   int           `json:"id"`
		Name string        `json:"name"`
		Tags []interface{} `json:"tags"`
	} `json:"consumer"`
	Contact              string    `json:"contact"`
	ContainsVmFilespaces bool      `json:"containsVmFilespaces"`
	CpuCount             int       `json:"cpuCount"`
	CreatedDate          time.Time `json:"createdDate"`
	DataSourceState      struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"dataSourceState"`
	DataSourceType struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"dataSourceType"`
	Decommissioned bool `json:"decommissioned"`
	Domain         struct {
		ArchiveRetention int       `json:"archiveRetention"`
		BackupRetention  int       `json:"backupRetention"`
		CreatedDate      time.Time `json:"createdDate"`
		Description      string    `json:"description"`
		ID               int       `json:"id"`
		MissingInTsm     bool      `json:"missingInTsm"`
		Name             string    `json:"name"`
		Server           struct {
			AllowedDedupStatThreads int           `json:"allowedDedupStatThreads"`
			ClientOptionSets        []interface{} `json:"clientOptionSets"`
			Consumers               []interface{} `json:"consumers"`
			CreatedDate             time.Time     `json:"createdDate"`
			DataCenter              struct {
				BusinessUnit struct {
					BillingStorageType               int           `json:"billingStorageType"`
					BillingStorageTypeID             int           `json:"billingStorageTypeId"`
					ClientOptionSetFilter            []interface{} `json:"clientOptionSetFilter"`
					Consumers                        []interface{} `json:"consumers"`
					CreatedDate                      time.Time     `json:"createdDate"`
					DomainFilter                     []interface{} `json:"domainFilter"`
					FinalDeleteRequestApprover       bool          `json:"finalDeleteRequestApprover"`
					ID                               int           `json:"id"`
					InvoiceDay                       int           `json:"invoiceDay"`
					Name                             string        `json:"name"`
					NodeLimit                        int           `json:"nodeLimit"`
					PasswordExpirationDays           int           `json:"passwordExpirationDays"`
					RegistrationNumber               string        `json:"registrationNumber"`
					ReportRemotely                   bool          `json:"reportRemotely"`
					RequireTotp                      bool          `json:"requireTotp"`
					RequiredApproversOfDeleteRequest int           `json:"requiredApproversOfDeleteRequest"`
					StorageLimit                     int           `json:"storageLimit"`
					SupportResponsible               bool          `json:"supportResponsible"`
					SystemID                         string        `json:"systemId"`
					Tags                             []interface{} `json:"tags"`
					TransferLimit                    int           `json:"transferLimit"`
					UseScheduleBindings              bool          `json:"useScheduleBindings"`
					UsersCanApproveOwnRequests       bool          `json:"usersCanApproveOwnRequests"`
				} `json:"businessUnit"`
				CreatedDate time.Time `json:"createdDate"`
				ID          int       `json:"id"`
				Name        string    `json:"name"`
			} `json:"dataCenter"`
			DedupStatsGenerationTime string `json:"dedupStatsGenerationTime"`
			DeleteMissingItems       bool   `json:"deleteMissingItems"`
			DeleteOldDedupStats      bool   `json:"deleteOldDedupStats"`
			Description              string `json:"description"`
			Domains                  []struct {
				ArchiveRetention int       `json:"archiveRetention"`
				BackupRetention  int       `json:"backupRetention"`
				CreatedDate      time.Time `json:"createdDate"`
				Description      string    `json:"description"`
				ID               int       `json:"id"`
				MissingInTsm     bool      `json:"missingInTsm"`
				Name             string    `json:"name"`
			} `json:"domains"`
			Errors                           []interface{} `json:"errors"`
			ErrorsDeactivated                bool          `json:"errorsDeactivated"`
			GenerateDedupStatsForStoragePool bool          `json:"generateDedupStatsForStoragePool"`
			GenerateDedupsStats              bool          `json:"generateDedupsStats"`
			GenerateDedupsStatsWaitSeconds   int           `json:"generateDedupsStatsWaitSeconds"`
			ID                               int           `json:"id"`
			ImportFromServer                 bool          `json:"importFromServer"`
			ImportOverlapMinutes             int           `json:"importOverlapMinutes"`
			ImportStoragePoolInfoFromScale   bool          `json:"importStoragePoolInfoFromScale"`
			InstallDate                      time.Time     `json:"installDate"`
			IsDefaultServerForBusinessUnit   bool          `json:"isDefaultServerForBusinessUnit"`
			LastSuccessfulImportTime         time.Time     `json:"lastSuccessfulImportTime"`
			LastSuccessfulReportTime         time.Time     `json:"lastSuccessfulReportTime"`
			Level                            int           `json:"level"`
			LowestAllowedClientVersion       string        `json:"lowestAllowedClientVersion"`
			LowestCompatibleClientVersion    string        `json:"lowestCompatibleClientVersion"`
			MachineName                      string        `json:"machineName"`
			Name                             string        `json:"name"`
			NodeNetworkAddress               string        `json:"nodeNetworkAddress"`
			NodeNetworkPort                  int           `json:"nodeNetworkPort"`
			Port                             int           `json:"port"`
			PortalNetworkAddress             string        `json:"portalNetworkAddress"`
			ReadOnly                         bool          `json:"readOnly"`
			Release                          int           `json:"release"`
			ReportingTime                    string        `json:"reportingTime"`
			SubLevel                         int           `json:"subLevel"`
			TimeZone                         struct {
				CreatedDate time.Time `json:"createdDate"`
				ID          int       `json:"id"`
				Name        string    `json:"name"`
				Offset      int       `json:"offset"`
				WindowsID   string    `json:"windowsId"`
			} `json:"timeZone"`
			Type struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"type"`
			UseSsl      bool   `json:"useSsl"`
			Version     int    `json:"version"`
			VersionName string `json:"versionName"`
		} `json:"server"`
	} `json:"domain"`
	Events          []interface{} `json:"events"`
	Filespaces      []interface{} `json:"filespaces"`
	Href            string        `json:"href"`
	Hypervisor      string        `json:"hypervisor"`
	ID              int           `json:"id"`
	LastAccessTime  time.Time     `json:"lastAccessTime"`
	Locked          bool          `json:"locked"`
	MacAddress      string        `json:"macAddress"`
	Name            string        `json:"name"`
	NodeGroups      []interface{} `json:"nodeGroups"`
	NodeSchedules   []interface{} `json:"nodeSchedules"`
	OperatingSystem struct {
		CreatedDate time.Time `json:"createdDate"`
		ID          int       `json:"id"`
		Name        string    `json:"name"`
		ShortName   string    `json:"shortName"`
	} `json:"operatingSystem"`
	OSLevel struct {
		CreatedDate time.Time `json:"createdDate"`
		ID          int       `json:"id"`
		Name        string    `json:"name"`
	} `json:"osLevel"`
	Platform struct {
		CreatedDate time.Time `json:"createdDate"`
		ID          int       `json:"id"`
		Name        string    `json:"name"`
	} `json:"platform"`
	ProxyAgents                 []interface{} `json:"proxyAgents"`
	ProxyTargets                []interface{} `json:"proxyTargets"`
	RegisteredTime              time.Time     `json:"registeredTime"`
	RemoteActivities            []interface{} `json:"remoteActivities"`
	ReplicationMode             int           `json:"replicationMode"`
	ReplicationServerPrimary    string        `json:"replicationServerPrimary"`
	ReplicationServerSecondary  string        `json:"replicationServerSecondary"`
	ReplicationServerSecondary2 string        `json:"replicationServerSecondary2"`
	ReplicationState            bool          `json:"replicationState"`
	RetentionSets               []interface{} `json:"retentionSets"`
	Schedules                   []struct {
		AsNodeNames   []interface{} `json:"asNodeNames"`
		CreatedDate   time.Time     `json:"createdDate"`
		DailyJobLists []interface{} `json:"dailyJobLists"`
		DayOfWeek     string        `json:"dayOfWeek"`
		Description   string        `json:"description"`
		Duration      int           `json:"duration"`
		Friday        bool          `json:"friday"`
		ID            int           `json:"id"`
		Monday        bool          `json:"monday"`
		Name          string        `json:"name"`
		Objects       string        `json:"objects"`
		Options       string        `json:"options"`
		Period        int           `json:"period"`
		Priority      int           `json:"priority"`
		ProxySchedule bool          `json:"proxySchedule"`
		Saturday      bool          `json:"saturday"`
		StartTime     time.Time     `json:"startTime"`
		Sunday        bool          `json:"sunday"`
		Thursday      bool          `json:"thursday"`
		Tuesday       bool          `json:"tuesday"`
		Wednesday     bool          `json:"wednesday"`
	} `json:"schedules"`
	SessionSecurity struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"sessionSecurity"`
	SyncState struct {
		CreatedDate time.Time `json:"createdDate"`
		ID          int       `json:"id"`
		Name        string    `json:"name"`
	} `json:"syncState"`
	TcpAddress       string `json:"tcpAddress"`
	TcpName          string `json:"tcpName"`
	Ticket           string `json:"ticket"`
	TsmClientVersion struct {
		CreatedDate time.Time `json:"createdDate"`
		ID          int       `json:"id"`
		Name        string    `json:"name"`
	} `json:"tsmClientVersion"`
	TsmName     string `json:"tsmName"`
	TsmNodeID   string `json:"tsmNodeId"`
	TsmPassword string `json:"tsmPassword"`
	Type        struct {
		CreatedDate time.Time `json:"createdDate"`
		ID          int       `json:"id"`
		Name        string    `json:"name"`
	} `json:"type"`
	UnscheduledActivities []interface{} `json:"unscheduledActivities"`
}
