package app

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/mint"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	terraappparams "github.com/classic-terra/core/app/params"

	customauth "github.com/classic-terra/core/custom/auth"
	customauthsim "github.com/classic-terra/core/custom/auth/simulation"
	customauthz "github.com/classic-terra/core/custom/authz"
	custombank "github.com/classic-terra/core/custom/bank"
	customcrisis "github.com/classic-terra/core/custom/crisis"
	customdistr "github.com/classic-terra/core/custom/distribution"
	customevidence "github.com/classic-terra/core/custom/evidence"
	customfeegrant "github.com/classic-terra/core/custom/feegrant"
	customgov "github.com/classic-terra/core/custom/gov"
	custommint "github.com/classic-terra/core/custom/mint"
	customparams "github.com/classic-terra/core/custom/params"
	customslashing "github.com/classic-terra/core/custom/slashing"
	customstaking "github.com/classic-terra/core/custom/staking"
	customupgrade "github.com/classic-terra/core/custom/upgrade"

	"github.com/classic-terra/core/x/market"
	markettypes "github.com/classic-terra/core/x/market/types"
	"github.com/classic-terra/core/x/oracle"
	oracletypes "github.com/classic-terra/core/x/oracle/types"
	"github.com/classic-terra/core/x/treasury"
	treasuryclient "github.com/classic-terra/core/x/treasury/client"
	treasurytypes "github.com/classic-terra/core/x/treasury/types"
	"github.com/classic-terra/core/x/vesting"
	"github.com/classic-terra/core/x/wasm"
	wasmtypes "github.com/classic-terra/core/x/wasm/types"

	// unnamed import of statik for swagger UI support
	_ "github.com/classic-terra/core/client/docs/statik"
)

var (
	// ModuleBasics = The ModuleBasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		customauth.AppModuleBasic{},
		customauthz.AppModuleBasic{},
		genutil.AppModuleBasic{},
		custombank.AppModuleBasic{},
		capability.AppModuleBasic{},
		customstaking.AppModuleBasic{},
		custommint.AppModuleBasic{},
		customdistr.AppModuleBasic{},
		customgov.NewAppModuleBasic(
			paramsclient.ProposalHandler,
			distrclient.ProposalHandler,
			upgradeclient.ProposalHandler,
			upgradeclient.CancelProposalHandler,
			treasuryclient.ProposalAddBurnTaxExemptionAddressHandler,
			treasuryclient.ProposalRemoveBurnTaxExemptionAddressHandler,
		),
		customparams.AppModuleBasic{},
		customcrisis.AppModuleBasic{},
		customslashing.AppModuleBasic{},
		customfeegrant.AppModuleBasic{},
		customupgrade.AppModuleBasic{},
		customevidence.AppModuleBasic{},
		vesting.AppModuleBasic{},
		oracle.AppModuleBasic{},
		market.AppModuleBasic{},
		treasury.AppModuleBasic{},
		wasm.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil, // just added to enable align fee
		treasurytypes.BurnModuleName:   {authtypes.Burner},
		minttypes.ModuleName:           {authtypes.Minter},
		markettypes.ModuleName:         {authtypes.Minter, authtypes.Burner},
		oracletypes.ModuleName:         nil,
		distrtypes.ModuleName:          nil,
		treasurytypes.ModuleName:       {authtypes.Minter, authtypes.Burner},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
	}

	// module accounts that are allowed to receive tokens
	allowedReceivingModAcc = map[string]bool{
		oracletypes.ModuleName:       true,
		treasurytypes.BurnModuleName: true,
	}
)

func appModules(
	app *TerraApp,
	encodingConfig terraappparams.EncodingConfig,
	skipGenesisInvariants bool,
) []module.AppModule {
	appCodec := encodingConfig.Marshaler

	return []module.AppModule{
		genutil.NewAppModule(
			app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.AccountKeeper, nil),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		params.NewAppModule(app.ParamsKeeper),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		market.NewAppModule(appCodec, app.MarketKeeper, app.AccountKeeper, app.BankKeeper, app.OracleKeeper),
		oracle.NewAppModule(appCodec, app.OracleKeeper, app.AccountKeeper, app.BankKeeper),
		treasury.NewAppModule(appCodec, app.TreasuryKeeper),
		wasm.NewAppModule(appCodec, app.WasmKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
	}
}

func simulationModules(
	app *TerraApp,
	encodingConfig terraappparams.EncodingConfig,
	_ bool,
) []module.AppModuleSimulation {
	appCodec := encodingConfig.Marshaler

	return []module.AppModuleSimulation{
		customauth.NewAppModule(appCodec, app.AccountKeeper, customauthsim.RandomGenesisAccounts),
		// TODO - uncomment when v0.43.0 fix the simulation bug
		// authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		custombank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		params.NewAppModule(app.ParamsKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		oracle.NewAppModule(appCodec, app.OracleKeeper, app.AccountKeeper, app.BankKeeper),
		market.NewAppModule(appCodec, app.MarketKeeper, app.AccountKeeper, app.BankKeeper, app.OracleKeeper),
		treasury.NewAppModule(appCodec, app.TreasuryKeeper),
		wasm.NewAppModule(appCodec, app.WasmKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
	}
}

func orderBeginBlockers() []string {
	return []string{
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		oracletypes.ModuleName,
		genutiltypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		treasurytypes.ModuleName,
		markettypes.ModuleName,
		wasmtypes.ModuleName,
	}
}

func orderEndBlockers() []string {
	return []string{
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		oracletypes.ModuleName,
		genutiltypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		treasurytypes.ModuleName,
		markettypes.ModuleName,
		wasmtypes.ModuleName,
	}
}

func orderInitGenesis() []string {
	return []string{
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		markettypes.ModuleName,
		oracletypes.ModuleName,
		treasurytypes.ModuleName,
		wasmtypes.ModuleName,
		authz.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		feegrant.ModuleName,
	}
}
