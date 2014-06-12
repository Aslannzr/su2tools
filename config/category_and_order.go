package config

// This file was autogenerated by github.com/btracey/su2tools/config/autogen_options/

type category struct {
	Id          int
	Name        string
	Description string
}

var categoryList = []category{
	{
		Id:          0,
		Name:        "Problem Definition",
		Description: "--- Options related to problem definition and partitioning ---",
	},
	{
		Id:          1,
		Name:        "Boundary Markers",
		Description: "--- Options related to various boundary markers ---",
	},
	{
		Id:          2,
		Name:        "Grid adaptation",
		Description: "--- Options related to grid adaptation ---",
	},
	{
		Id:          3,
		Name:        "Time-marching",
		Description: "--- Options related to time-marching ---",
	},
	{
		Id:          4,
		Name:        "Linear solver definition",
		Description: "--- Options related to the linear solvers ---",
	},
	{
		Id:          5,
		Name:        "Dynamic mesh definition",
		Description: "--- Options related to dynamic meshes ---",
	},
	{
		Id:          6,
		Name:        "Wind Gust",
		Description: "--- Options related to wind gust simulations ---",
	},
	{
		Id:          7,
		Name:        "Convergence",
		Description: "--- Options related to convergence ---",
	},
	{
		Id:          8,
		Name:        "Multi-grid",
		Description: "--- Options related to Multi-grid ---",
	},
	{
		Id:          9,
		Name:        "Spatial Discretization",
		Description: "--- Options related to the spatial discretization ---",
	},
	{
		Id:          10,
		Name:        "Adjoint and Gradient",
		Description: "--- Options related to the adjoint and gradient ---",
	},
	{
		Id:          11,
		Name:        "Input/output files and formats",
		Description: "--- Options related to input/output files and formats ---",
	},
	{
		Id:          12,
		Name:        "Equivalent Area",
		Description: "--- Options related to the equivalent area ---",
	},
	{
		Id:          13,
		Name:        "Freestream Conditions",
		Description: "--- Options related to freestream specification ---",
	},
	{
		Id:          14,
		Name:        "Reference Conditions",
		Description: "--- Options related to reference values for nondimensionalization ---",
	},
	{
		Id:          15,
		Name:        "Reacting Flow",
		Description: "--- Options related to the reacting gas mixtures ---",
	},
	{
		Id:          16,
		Name:        "Free surface simulation",
		Description: "--- Options related to free surface simulation ---",
	},
	{
		Id:          17,
		Name:        "Grid deformation",
		Description: "--- Options related to the grid deformation ---",
	},
	{
		Id:          18,
		Name:        "Rotorcraft problem",
		Description: "--- option related to rotorcraft problems ---",
	},
	{
		Id:          19,
		Name:        "FEA solver",
		Description: "--- Options related to the FEA solver ---",
	},
	{
		Id:          20,
		Name:        "Wave solver",
		Description: "--- options related to the wave solver ---",
	},
	{
		Id:          21,
		Name:        "Heat solver",
		Description: "--- options related to the heat solver ---",
	},
	{
		Id:          22,
		Name:        "Visualize Control Volumes",
		Description: "--- options related to visualizing control volumes ---",
	},
	{
		Id:          23,
		Name:        "Inverse design problem",
		Description: "--- options related to inverse design problem ---",
	},
	{
		Id:          24,
		Name:        "Unsupported options",
		Description: "--- Options that are experimental and not intended for general use ---",
	},
	{
		Id:          25,
		Name:        "FFD point inversion",
		Description: " DESCRIPTION: Number of total iterations in the FFD point inversion ",
	},
}

// optionList contains a list of the options to be printed. Outer slice is
// for the category, inner list is the order of the options
var optionList = [][]Option{
	{
		RegimeType,
		ExtraOutput,
		PhysicalProblem,
		MathProblem,
		KindTurbModel,
		KindTransModel,
		Axisymmetric,
		GravityForce,
		LowFidelitySimulation,
		RestartSol,
		VisualizePart,
	},
	{
		MarkerPlotting,
		MarkerMonitoring,
		MarkerDesigning,
		GeoMarker,
		MarkerEuler,
		MarkerFar,
		MarkerSym,
		MarkerPressure,
		MarkerNearfield,
		MarkerInterface,
		MarkerDirichlet,
		MarkerNeumann,
		ElecDirichlet,
		ElecNeumann,
		MarkerCustom,
		MarkerPeriodic,
		MarkerActdisk,
		InletType,
		MarkerInlet,
		MarkerSupersonicInlet,
		MarkerOutlet,
		MarkerIsothermal,
		MarkerIsothermalNoncatalytic,
		MarkerIsothermalCatalytic,
		MarkerHeatflux,
		MarkerHeatfluxNoncatalytic,
		MarkerHeatfluxCatalytic,
		MarkerNacelleInflow,
		SubsonicNacelleInflow,
		MarkerNacelleExhaust,
		MarkerNormalDispl,
		MarkerNormalLoad,
		MarkerFlowload,
		DampNacelleInflow,
		MarkerOut1d,
	},
	{
		KindAdapt,
		NewElems,
		DualvolPower,
		AnalyticalSurfdef,
		SmoothGeometry,
		AdaptBoundary,
		DivideElements,
	},
	{
		UnsteadySimulation,
		CflNumber,
		CflRamp,
		CflReductionAdjflow,
		CflReductionTurb,
		CflReductionAdjturb,
		ExtIter,
		RkAlphaCoeff,
		UnstTimestep,
		UnstTime,
		UnstCflNumber,
		UnstIntIter,
		TimeInstances,
		UnstRestartIter,
		UnstAdjointIter,
		TimeDiscreFlow,
		TimeDiscreTne2,
		TimeDiscreAdjtne2,
		TimeDiscreAdjlevelset,
		TimeDiscreAdjflow,
		TimeDiscreLin,
		TimeDiscreTurb,
		TimeDiscreAdjturb,
		TimeDiscreWave,
		TimeDiscreFea,
		TimeDiscreHeat,
		TimeDiscrePoisson,
	},
	{
		LinearSolver,
		LinearSolverPrec,
		LinearSolverError,
		LinearSolverIter,
		LinearSolverRestartFrequency,
		LinearSolverRelax,
		RoeTurkelPrec,
		MinRoeTurkelPrec,
		MaxRoeTurkelPrec,
		AdjturbLinSolver,
		AdjturbLinPrec,
		AdjturbLinError,
		AdjturbLinIter,
	},
	{
		GridMovement,
		GridMovementKind,
		MarkerMoving,
		MachMotion,
		MotionOriginX,
		MotionOriginY,
		MotionOriginZ,
		TranslationRateX,
		TranslationRateY,
		TranslationRateZ,
		RotationRateX,
		RotationRateY,
		RotationRateZ,
		PitchingOmegaX,
		PitchingOmegaY,
		PitchingOmegaZ,
		PitchingAmplX,
		PitchingAmplY,
		PitchingAmplZ,
		PitchingPhaseX,
		PitchingPhaseY,
		PitchingPhaseZ,
		PlungingOmegaX,
		PlungingOmegaY,
		PlungingOmegaZ,
		PlungingAmplX,
		PlungingAmplY,
		PlungingAmplZ,
		MoveMotionOrigin,
		MotionFilename,
		FreqPlungeAeroelastic,
		FreqPitchAeroelastic,
	},
	{
		WindGust,
		GustType,
		GustWavelength,
		GustPeriods,
		GustAmpl,
		GustBeginTime,
		GustBeginLoc,
		GustDir,
	},
	{
		ConvCriteria,
		ResidualReduction,
		ResidualMinval,
		StartconvIter,
		CauchyElems,
		CauchyEps,
		CauchyFuncFlow,
		CauchyFuncAdjflow,
		CauchyFuncLin,
		FullmgCauchyEps,
	},
	{
		Fullmg,
		StartUpIter,
		Mglevel,
		Mgcycle,
		MgPreSmooth,
		MgPostSmooth,
		MgCorrectionSmooth,
		MgDampRestriction,
		MgDampProlongation,
		MgCflReduction,
		MaxChildren,
		MaxDimension,
	},
	{
		NumMethodGrad,
		LimiterCoeff,
		SharpEdgesCoeff,
		ConvNumMethodFlow,
		ViscNumMethodFlow,
		SourNumMethodFlow,
		SpatialOrderFlow,
		SlopeLimiterFlow,
		AdCoeffFlow,
		ConvNumMethodAdjflow,
		ViscNumMethodAdjflow,
		SourNumMethodAdjflow,
		SpatialOrderAdjflow,
		SlopeLimiterAdjflow,
		AdCoeffAdjflow,
		SpatialOrderTurb,
		SlopeLimiterTurb,
		ConvNumMethodTurb,
		ViscNumMethodTurb,
		SourNumMethodTurb,
		SpatialOrderAdjturb,
		SlopeLimiterAdjturb,
		ConvNumMethodAdjturb,
		ViscNumMethodAdjturb,
		SourNumMethodAdjturb,
		ConvNumMethodLin,
		ViscNumMethodLin,
		SourNumMethodLin,
		AdCoeffLin,
		SpatialOrderAdjlevelset,
		SlopeLimiterAdjlevelset,
		ConvNumMethodAdjlevelset,
		ViscNumMethodAdjlevelset,
		SourNumMethodAdjlevelset,
		ConvNumMethodTne2,
		ViscNumMethodTne2,
		SourNumMethodTne2,
		SpatialOrderTne2,
		SlopeLimiterTne2,
		AdCoeffTne2,
		ConvNumMethodAdjtne2,
		ViscNumMethodAdjtne2,
		SourNumMethodAdjtne2,
		SpatialOrderAdjtne2,
		SlopeLimiterAdjtne2,
		AdCoeffAdjtne2,
		ViscNumMethodWave,
		SourNumMethodWave,
		ViscNumMethodPoisson,
		SourNumMethodPoisson,
		ViscNumMethodFea,
		SourNumMethodFea,
		ViscNumMethodHeat,
		SourNumMethodHeat,
		SourNumMethodTemplate,
	},
	{
		LimitAdjflow,
		ObjectiveFunction,
		GeoLocationSections,
		GeoOrientationSections,
		GeoNumberSections,
		GeoVolumeSections,
		GeoPlotSections,
		GeoMode,
		DragInSonicboom,
		SensSmoothing,
		ContinuousEqns,
		DiscreteEqns,
		FrozenVisc,
		CteViscousDrag,
		FixAzimuthalLine,
		SensRemoveSharp,
		PnormHeatflux,
	},
	{
		OutputFormat,
		MeshFormat,
		CgnsToSu2,
		MeshFilename,
		MeshScaleChange,
		MeshOutput,
		MeshOutFilename,
		ConvFilename,
		SolutionFlowFilename,
		SolutionLinFilename,
		SolutionAdjFilename,
		RestartFlowFilename,
		RestartLinFilename,
		RestartAdjFilename,
		RestartWaveFilename,
		VolumeFlowFilename,
		VolumeStructureFilename,
		SurfaceStructureFilename,
		SurfaceWaveFilename,
		SurfaceHeatFilename,
		VolumeWaveFilename,
		VolumeHeatFilename,
		VolumeAdjwaveFilename,
		VolumeAdjFilename,
		VolumeLinFilename,
		GradObjfuncFilename,
		ValueObjfuncFilename,
		SurfaceFlowFilename,
		SurfaceAdjFilename,
		SurfaceLinFilename,
		WrtSolFreq,
		WrtSolFreqDualtime,
		WrtConFreq,
		WrtConFreqDualtime,
		WrtVolSol,
		WrtSrfSol,
		WrtCsvSol,
		WrtRestart,
		WrtResiduals,
		WrtHalo,
		Wrt1dOutput,
	},
	{
		EquivArea,
		EaIntLimit,
	},
	{
		GasConstant,
		GammaValue,
		ReynoldsNumber,
		ReynoldsLength,
		PrandtlLam,
		PrandtlTurb,
		BulkModulus,
		ArtcompFactor,
		MachNumber,
		FreestreamPressure,
		FreestreamDensity,
		FreestreamTemperature,
		FreestreamTemperatureVe,
		FreestreamVelocity,
		FreestreamViscosity,
		FreestreamIntermittency,
		FreestreamTurbulenceintensity,
		FreestreamNuFactor,
		FreestreamTurb2lamviscratio,
		SideslipAngle,
		Aoa,
		FixedClMode,
		TargetCl,
		DampFixedCl,
	},
	{
		RefOriginMomentX,
		RefOriginMomentY,
		RefOriginMomentZ,
		RefArea,
		RefLengthMoment,
		RefElemLength,
		RefSharpEdges,
		RefPressure,
		RefTemperature,
		RefDensity,
		RefVelocity,
		RefViscosity,
	},
	{
		GasModel,
		GasComposition,
	},
	{
		RatioDensity,
		RatioViscosity,
		FreesurfaceZero,
		FreesurfaceDepth,
		FreesurfaceThickness,
		FreesurfaceDampingCoeff,
		FreesurfaceDampingLength,
		FreesurfaceOutlet,
	},
	{
		DvKind,
		DvMarker,
		DvValue,
		DvParam,
		HoldGridFixed,
		HoldGridFixedCoord,
		VisualizeDeformation,
		DeformConsoleOutput,
		DeformNonlinearIter,
		DeformLinearIter,
		DeformTolFactor,
		DeformStiffnessType,
		YoungsModulus,
		PoissonsRatio,
	},
	{
		CyclicPitch,
		CollectivePitch,
	},
	{
		ElasticityModulus,
		PoissonRatio,
		MaterialDensity,
	},
	{
		WaveSpeed,
	},
	{
		ThermalDiffusivity,
	},
	{
		VisualizeCv,
	},
	{
		InvDesignCp,
		InvDesignHeatflux,
	},
	{
		MlTurbModelFile,
		MlTurbModelFeatureset,
		MlTurbModelExtra,
	},
	{
		FfdIterations,
		FfdTolerance,
	},
}
