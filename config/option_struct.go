package config

// This file was autogenerated by github.com/btracey/su2tools/config/autogen_options/

import (
	"github.com/btracey/su2tools/config/enum"
	"github.com/btracey/su2tools/config/su2types"
)

type Options struct {
	// Adjoint type
	RegimeType enum.Regime
	// Physical governing equations
	PhysicalProblem enum.Solver
	// Mathematical problem
	MathProblem enum.MathProblem
	// Specify turbulence model
	KindTurbModel enum.TurbModel
	// Specify transition model
	KindTransModel enum.TransModel
	// Axisymmetric simulation
	Axisymmetric bool
	// Add the gravity force
	GravityForce bool
	// Perform a low fidelity simulation
	LowFidelitySimulation bool
	// Restart solution from native solution file
	RestartSol bool
	// Write a tecplot file for each partition
	VisualizePart bool
	// System of measurements
	SystemMeasurements enum.Measurements
	// Fluid model
	FluidModel enum.Fluidmodel
	// Specific gas constant (287.058 J/kg*K (air), only for compressible flows)
	GasConstant float64
	// Ratio of specific heats (1.4 (air), only for compressible flows)
	GammaValue float64
	// Critical Temperature, default value for AIR
	CriticalTemperature float64
	// Critical Pressure, default value for MDM
	CriticalPressure float64
	// Critical Density, default value for MDM
	CriticalDensity float64
	// Critical Density, default value for MDM
	AcentricFactor float64
	// model of the viscosity
	ViscosityModel enum.Viscositymodel
	// Critical Temperature, default value for AIR
	MuConstant float64
	// Sutherland Viscosity Ref default value for AIR SI
	MuRef float64
	// Sutherland Temperature Ref, default value for AIR SI
	MuTRef float64
	// Sutherland constant, default value for AIR SI
	SutherlandConstant float64
	// Reynolds number (non-dimensional, based on the free-stream values)
	ReynoldsNumber float64
	// Reynolds length (1 m by default)
	ReynoldsLength float64
	// Laminar Prandtl number (0.72 (air), only for compressible flows)
	PrandtlLam float64
	// Turbulent Prandtl number (0.9 (air), only for compressible flows)
	PrandtlTurb float64
	// Value of the Bulk Modulus
	BulkModulus float64
	// Artifical compressibility factor
	ArtcompFactor float64
	// Mach number (non-dimensional, based on the free-stream values)
	MachNumber float64
	// Free-stream option to choose between density and temperature for initializing the solution
	FreestreamOption enum.FreestreamOption
	// Free-stream pressure (101325.0 N/m^2 by default)
	FreestreamPressure float64
	// Free-stream density (1.2886 Kg/m^3 (air), 998.2 Kg/m^3 (water))
	FreestreamDensity float64
	// Free-stream temperature (288.15 K by default)
	FreestreamTemperature float64
	// Free-stream vibrational-electronic temperature (288.15 K by default)
	FreestreamTemperatureVe float64
	// Free-stream velocity (m/s)
	FreestreamVelocity [3]float64
	// Free-stream viscosity (1.853E-5 Ns/m^2 (air), 0.798E-3 Ns/m^2 (water))
	FreestreamViscosity float64
	//
	FreestreamIntermittency float64
	//
	FreestreamTurbulenceintensity float64
	//
	FreestreamNuFactor float64
	//
	FreestreamTurb2lamviscratio float64
	// Side-slip angle (degrees, only for compressible flows)
	SideslipAngle float64
	// Angle of attack (degrees, only for compressible flows)
	Aoa float64
	// Activate fixed CL mode (specify a CL instead of AoA).
	FixedClMode bool
	// Specify a fixed coefficient of lift instead of AoA (only for compressible flows)
	TargetCl float64
	// Damping factor for fixed CL mode.
	DampFixedCl float64
	// X Reference origin for moment computation
	RefOriginMomentX []float64
	// Y Reference origin for moment computation
	RefOriginMomentY []float64
	// Z Reference origin for moment computation
	RefOriginMomentZ []float64
	// Reference area for force coefficients (0 implies automatic calculation)
	RefArea float64
	// Reference length for pitching, rolling, and yawing non-dimensional moment
	RefLengthMoment float64
	// Reference element length for computing the slope limiter epsilon
	RefElemLength float64
	// Reference coefficient for detecting sharp edges
	RefSharpEdges float64
	// Reference pressure (1.0 N/m^2 by default, only for compressible flows)
	RefPressure float64
	// Reference temperature (1.0 K by default, only for compressible flows)
	RefTemperature float64
	// Reference density (1.0 Kg/m^3 by default, only for compressible flows)
	RefDensity float64
	// Reference velocity (incompressible only)
	RefVelocity float64
	// Reference viscosity (incompressible only)
	RefViscosity float64
	// Marker(s) of the surface in the surface flow solution file
	MarkerPlotting []string
	// Marker(s) of the surface where evaluate the non-dimensional coefficients
	MarkerMonitoring []string
	// Marker(s) of the surface where objective function (design problem) will be evaluated
	MarkerDesigning []string
	// Marker(s) of the surface where evaluate the geometrical functions
	GeoMarker []string
	// Euler wall boundary marker(s)
	MarkerEuler []string
	// Far-field boundary marker(s)
	MarkerFar []string
	// Symmetry boundary condition
	MarkerSym []string
	// Symmetry boundary condition
	MarkerPressure []string
	// Near-Field boundary condition
	MarkerNearfield []string
	// Zone interface boundary marker(s)
	MarkerInterface []string
	// Dirichlet boundary marker(s)
	MarkerDirichlet []string
	// Neumann boundary marker(s)
	MarkerNeumann []string
	// poisson dirichlet boundary marker(s)
	ElecDirichlet *su2types.StringDoubleList
	// poisson neumann boundary marker(s)
	ElecNeumann []string
	// Custom boundary marker(s)
	MarkerCustom []string
	// Periodic boundary marker(s) for use with SU2_MSHFormat: ( periodic marker, donor marker, rotation_center_x, rotation_center_y,rotation_center_z, rotation_angle_x-axis, rotation_angle_y-axis,rotation_angle_z-axis, translation_x, translation_y, translation_z, ... )
	MarkerPeriodic *su2types.Periodic
	// Periodic boundary marker(s) for use with SU2_MSHFormat: ( periodic marker, donor marker, rotation_center_x, rotation_center_y,rotation_center_z, rotation_angle_x-axis, rotation_angle_y-axis,rotation_angle_z-axis, translation_x, translation_y, translation_z, ... )
	MarkerActdisk *su2types.ActuatorDisk
	// Inlet boundary type
	InletType enum.InletType
	// Inlet boundary marker(s) with the following formats,Total Conditions: (inlet marker, total temp, total pressure, flow_direction_x,flow_direction_y, flow_direction_z, ... ) where flow_direction isa unit vector.Mass Flow: (inlet marker, density, velocity magnitude, flow_direction_x,flow_direction_y, flow_direction_z, ... ) where flow_direction isa unit vector.
	MarkerInlet *su2types.Inlet
	// Riemann boundary marker(s) with the following formats, a unit vector.
	MarkerRiemann *su2types.Riemann
	// % Supersonic inlet boundary marker(s)Format: (inlet marker, temperature, static pressure, velocity_x,velocity_y, velocity_z, ... ), i.e. primitive variables specified.
	MarkerSupersonicInlet *su2types.Inlet
	// Outlet boundary marker(s)Format: ( outlet marker, back pressure (static), ... )
	MarkerOutlet *su2types.StringDoubleList
	// Isothermal wall boundary marker(s)Format: ( isothermal marker, wall temperature (static), ... )
	MarkerIsothermal *su2types.StringDoubleList
	// Isothermal wall boundary marker(s)Format: ( isothermal marker, wall temperature (static), ... )
	MarkerIsothermalNoncatalytic *su2types.StringDoubleList
	// Isothermal wall boundary marker(s)Format: ( isothermal marker, wall temperature (static), ... )
	MarkerIsothermalCatalytic *su2types.StringDoubleList
	// Specified heat flux wall boundary marker(s)Format: ( Heat flux marker, wall heat flux (static), ... )
	MarkerHeatflux *su2types.StringDoubleList
	// Specified heat flux wall boundary marker(s)Format: ( Heat flux marker, wall heat flux (static), ... )
	MarkerHeatfluxNoncatalytic *su2types.StringDoubleList
	// Specified heat flux wall boundary marker(s)Format: ( Heat flux marker, wall heat flux (static), ... )
	MarkerHeatfluxCatalytic *su2types.StringDoubleList
	// Nacelle inflow boundary marker(s)Format: ( nacelle inflow marker, fan face Mach, ... )
	MarkerNacelleInflow *su2types.StringDoubleList
	// Engine subsonic intake region
	SubsonicNacelleInflow bool
	// Nacelle exhaust boundary marker(s)Format: (nacelle exhaust marker, total nozzle temp, total nozzle pressure, ... )
	MarkerNacelleExhaust *su2types.InletFixed
	// Displacement boundary marker(s)
	MarkerNormalDispl *su2types.StringDoubleList
	// Load boundary marker(s)
	MarkerNormalLoad *su2types.StringDoubleList
	// Flow load boundary marker(s)
	MarkerFlowload *su2types.StringDoubleList
	// Damping factor for engine inlet condition
	DampNacelleInflow float64
	// Outlet boundary marker(s) over which to calculate 1-D flow propertiesFormat: ( outlet marker)
	MarkerOut1d []string
	// Unsteady simulation
	UnsteadySimulation enum.Unsteady
	// Courant-Friedrichs-Lewy condition of the finest grid
	CflNumber float64
	// CFL ramp (factor, number of iterations, CFL limit)
	CflRamp [3]float64
	// Reduction factor of the CFL coefficient in the adjoint problem
	CflReductionAdjflow float64
	// Reduction factor of the CFL coefficient in the level set problem
	CflReductionTurb float64
	// Reduction factor of the CFL coefficient in the turbulent adjoint problem
	CflReductionAdjturb float64
	// Number of total iterations
	ExtIter uint64
	// Runge-Kutta alpha coefficients
	RkAlphaCoeff []float64
	// Time Step for dual time stepping simulations (s)
	UnstTimestep float64
	// Total Physical Time for dual time stepping simulations (s)
	UnstTime float64
	// Unsteady Courant-Friedrichs-Lewy number of the finest grid
	UnstCflNumber float64
	// Number of internal iterations (dual time method)
	UnstIntIter uint64
	// Integer number of periodic time instances for Time Spectral
	TimeInstances uint16
	// Iteration number to begin unsteady restarts (dual time method)
	UnstRestartIter int
	// Starting direct solver iteration for the unsteady adjoint
	UnstAdjointIter int
	// Time discretization
	TimeDiscreFlow enum.TimeInt
	// Time discretization
	TimeDiscreTne2 enum.TimeInt
	// Time discretization
	TimeDiscreAdjtne2 enum.TimeInt
	// Time discretization
	TimeDiscreAdjlevelset enum.TimeInt
	// Time discretization
	TimeDiscreAdjflow enum.TimeInt
	// Time discretization
	TimeDiscreLin enum.TimeInt
	// Time discretization
	TimeDiscreTurb enum.TimeInt
	// Time discretization
	TimeDiscreAdjturb enum.TimeInt
	// Time discretization
	TimeDiscreWave enum.TimeInt
	// Time discretization
	TimeDiscreFea enum.TimeInt
	// Time discretization
	TimeDiscreHeat enum.TimeInt
	// Time discretization
	TimeDiscrePoisson enum.TimeInt
	// Linear solver for the implicit, mesh deformation, or discrete adjoint systems
	LinearSolver enum.LinearSolver
	// Preconditioner for the Krylov linear solvers
	LinearSolverPrec enum.LinearSolverPrec
	// Minimum error threshold for the linear solver for the implicit formulation
	LinearSolverError float64
	// Maximum number of iterations of the linear solver for the implicit formulation
	LinearSolverIter uint64
	// Maximum number of iterations of the linear solver for the implicit formulation
	LinearSolverRestartFrequency uint64
	// Relaxation of the linear solver for the implicit formulation
	LinearSolverRelax float64
	// Roe-Turkel preconditioning for low Mach number flows
	RoeTurkelPrec bool
	// Time Step for dual time stepping simulations (s)
	MinRoeTurkelPrec float64
	// Time Step for dual time stepping simulations (s)
	MaxRoeTurkelPrec float64
	// Linear solver for the turbulent adjoint systems
	AdjturbLinSolver enum.LinearSolver
	// Preconditioner for the turbulent adjoint Krylov linear solvers
	AdjturbLinPrec enum.LinearSolverPrec
	// Minimum error threshold for the turbulent adjoint linear solver for the implicit formulation
	AdjturbLinError float64
	// Maximum number of iterations of the turbulent adjoint linear solver for the implicit formulation
	AdjturbLinIter uint16
	// Entropy fix factor
	EntropyFixCoeff float64
	// Convergence criteria
	ConvCriteria enum.ConvergeCrit
	// Residual reduction (order of magnitude with respect to the initial value)
	ResidualReduction float64
	// Min value of the residual (log10 of the residual)
	ResidualMinval float64
	// Iteration number to begin convergence monitoring
	StartconvIter uint64
	// Number of elements to apply the criteria
	CauchyElems uint16
	// Epsilon to control the series convergence
	CauchyEps float64
	// Flow functional for the Cauchy criteria
	CauchyFuncFlow enum.Objective
	// Adjoint functional for the Cauchy criteria
	CauchyFuncAdjflow enum.Sens
	// Linearized functional for the Cauchy criteria
	CauchyFuncLin enum.LinearObj
	// Epsilon for a full multigrid method evaluation
	FullmgCauchyEps float64
	// Full multi-grid
	Fullmg bool
	// Start up iterations using the fine grid only
	StartUpIter uint16
	// Multi-grid Levels
	Mglevel uint16
	// Multi-grid Cycle (0 = V cycle, 1 = W Cycle)
	Mgcycle uint16
	// Multi-grid pre-smoothing level
	MgPreSmooth []uint16
	// Multi-grid post-smoothing level
	MgPostSmooth []uint16
	// Jacobi implicit smoothing of the correction
	MgCorrectionSmooth []uint16
	// Damping factor for the residual restriction
	MgDampRestriction float64
	// Damping factor for the correction prolongation
	MgDampProlongation float64
	// Maximum number of children in the agglomeration stage
	MaxChildren uint16
	// Maximum length of an agglomerated element (relative to the domain)
	MaxDimension float64
	// Numerical method for spatial gradients
	NumMethodGrad enum.FlowGradient
	// Coefficient for the limiter
	LimiterCoeff float64
	// Freeze the value of the limiter after a number of iterations
	LimiterIter uint64
	// Coefficient for detecting the limit of the sharp edges
	SharpEdgesCoeff float64
	// Convective numerical method
	ConvNumMethodFlow string
	// Spatial numerical order integration
	SpatialOrderFlow enum.SpatialOrder
	// Slope limiter
	SlopeLimiterFlow enum.Limiter
	// 1st, 2nd and 4th order artificial dissipation coefficients
	AdCoeffFlow [3]float64
	// Convective numerical method
	ConvNumMethodAdjflow string
	// Spatial numerical order integration
	SpatialOrderAdjflow enum.SpatialOrder
	// Slope limiter
	SlopeLimiterAdjflow enum.Limiter
	// 1st, 2nd and 4th order artificial dissipation coefficients
	AdCoeffAdjflow [3]float64
	// Spatial numerical order integration
	SpatialOrderTurb enum.SpatialOrder
	// Slope limiter
	SlopeLimiterTurb enum.Limiter
	// Convective numerical method
	ConvNumMethodTurb string
	// Spatial numerical order integration
	SpatialOrderAdjturb enum.SpatialOrder
	// Slope limiter
	SlopeLimiterAdjturb enum.Limiter
	// Convective numerical method
	ConvNumMethodAdjturb string
	// Convective numerical method
	ConvNumMethodLin string
	// 1st, 2nd and 4th order artificial dissipation coefficients
	AdCoeffLin [2]float64
	// Spatial numerical order integration
	SpatialOrderAdjlevelset enum.SpatialOrder
	// Slope limiter
	SlopeLimiterAdjlevelset enum.Limiter
	// Convective numerical method
	ConvNumMethodAdjlevelset string
	// Convective numerical method
	ConvNumMethodTne2 string
	// Spatial numerical order integration
	SpatialOrderTne2 enum.SpatialOrder
	// Slope limiter
	SlopeLimiterTne2 enum.Limiter
	// 1st, 2nd and 4th order artificial dissipation coefficients
	AdCoeffTne2 [3]float64
	// Convective numerical method
	ConvNumMethodAdjtne2 string
	// Spatial numerical order integration
	SpatialOrderAdjtne2 enum.SpatialOrder
	// Slope limiter
	SlopeLimiterAdjtne2 enum.Limiter
	// 1st, 2nd and 4th order artificial dissipation coefficients
	AdCoeffAdjtne2 [3]float64
	// Limit value for the adjoint variable
	LimitAdjflow float64
	// Adjoint problem boundary condition
	ObjectiveFunction enum.Objective
	// Definition of the airfoil section
	GeoLocationSections [2]float64
	// Identify the axis of the section
	GeoOrientationSections enum.AxisOrientation
	// Percentage of new elements (% of the original number of elements)
	GeoNumberSections uint16
	// Number of section cuts to make when calculating internal volume
	GeoVolumeSections uint16
	// Output sectional forces for specified markers.
	GeoPlotSections bool
	// Mode of the GDC code (analysis, or gradient)
	GeoMode enum.GeometryMode
	// Drag weight in sonic boom Objective Function (from 0.0 to 1.0)
	DragInSonicboom float64
	// Sensitivity smoothing
	SensSmoothing enum.SensSmoothing
	// Continuous governing equation set
	ContinuousEqns enum.ContinuousEqns
	// Discrete governing equation set
	DiscreteEqns enum.DiscreteEqns
	// Adjoint frozen viscosity
	FrozenVisc bool
	//
	CteViscousDrag float64
	//
	FixAzimuthalLine float64
	// Remove sharp edges from the sensitivity evaluation
	SensRemoveSharp bool
	// P-norm for heat-flux based objective functions.
	PnormHeatflux float64
	// I/O
	OutputFormat enum.Output
	// Mesh input file format
	MeshFormat enum.Input
	// Convert a CGNS mesh to SU2 format
	CgnsToSu2 bool
	// Mesh input file
	MeshFilename string
	// Factor for scaling the mesh
	MeshScaleChange float64
	// Write a new mesh converted to meters
	MeshOutput bool
	// Cuthill–McKee ordering algorithm
	CuthillMckeeOrdering bool
	// Mesh output file
	MeshOutFilename string
	// Output file convergence history (w/o extension)
	ConvFilename string
	// Restart flow input file
	SolutionFlowFilename string
	// Restart linear flow input file
	SolutionLinFilename string
	// Restart adjoint input file
	SolutionAdjFilename string
	// Output file restart flow
	RestartFlowFilename string
	// Output file linear flow
	RestartLinFilename string
	// Output file restart adjoint
	RestartAdjFilename string
	// Output file restart wave
	RestartWaveFilename string
	// Output file flow (w/o extension) variables
	VolumeFlowFilename string
	// Output file structure (w/o extension) variables
	VolumeStructureFilename string
	// Output file structure (w/o extension) variables
	SurfaceStructureFilename string
	// Output file structure (w/o extension) variables
	SurfaceWaveFilename string
	// Output file structure (w/o extension) variables
	SurfaceHeatFilename string
	// Output file wave (w/o extension) variables
	VolumeWaveFilename string
	// Output file wave (w/o extension) variables
	VolumeHeatFilename string
	// Output file adj. wave (w/o extension) variables
	VolumeAdjwaveFilename string
	// Output file adjoint (w/o extension) variables
	VolumeAdjFilename string
	// Output file linear (w/o extension) variables
	VolumeLinFilename string
	// Output objective function gradient
	GradObjfuncFilename string
	// Output objective function
	ValueObjfuncFilename string
	// Output file surface flow coefficient (w/o extension)
	SurfaceFlowFilename string
	// Output file surface adjoint coefficient (w/o extension)
	SurfaceAdjFilename string
	// Output file surface linear coefficient (w/o extension)
	SurfaceLinFilename string
	// Writing solution file frequency
	WrtSolFreq uint64
	// Writing solution file frequency
	WrtSolFreqDualtime uint64
	// Writing convergence history frequency
	WrtConFreq uint64
	// Writing convergence history frequency for the dual time
	WrtConFreqDualtime uint64
	// Write a volume solution file
	WrtVolSol bool
	// Write a surface solution file
	WrtSrfSol bool
	// Write a surface CSV solution file
	WrtCsvSol bool
	// Write a restart solution file
	WrtRestart bool
	// Output residual info to solution/restart file
	WrtResiduals bool
	// Output residual info to solution/restart file
	WrtLimiters bool
	// Output residual info to solution/restart file
	WrtSharpedges bool
	// Output the rind layers in the solution files
	WrtHalo bool
	// Output averaged stagnation pressure on specified exit marker.
	OneDOutput bool
	// Mesh motion for unsteady simulations
	GridMovement bool
	// Type of mesh motion
	GridMovementKind []enum.Gridmovement
	// Marker(s) of moving surfaces (MOVING_WALL or DEFORMING grid motion).
	MarkerMoving []string
	// Mach number (non-dimensional, based on the mesh velocity and freestream vals.)
	MachMotion float64
	// Coordinates of the rigid motion origin
	MotionOriginX []float64
	// Coordinates of the rigid motion origin
	MotionOriginY []float64
	// Coordinates of the rigid motion origin
	MotionOriginZ []float64
	// Translational velocity vector (m/s) in the x, y, & z directions (RIGID_MOTION only)
	TranslationRateX []float64
	// Translational velocity vector (m/s) in the x, y, & z directions (RIGID_MOTION only)
	TranslationRateY []float64
	// Translational velocity vector (m/s) in the x, y, & z directions (RIGID_MOTION only)
	TranslationRateZ []float64
	// Angular velocity vector (rad/s) about x, y, & z axes (RIGID_MOTION only)
	RotationRateX []float64
	// Angular velocity vector (rad/s) about x, y, & z axes (RIGID_MOTION only)
	RotationRateY []float64
	// Angular velocity vector (rad/s) about x, y, & z axes (RIGID_MOTION only)
	RotationRateZ []float64
	// Pitching angular freq. (rad/s) about x, y, & z axes (RIGID_MOTION only)
	PitchingOmegaX []float64
	// Pitching angular freq. (rad/s) about x, y, & z axes (RIGID_MOTION only)
	PitchingOmegaY []float64
	// Pitching angular freq. (rad/s) about x, y, & z axes (RIGID_MOTION only)
	PitchingOmegaZ []float64
	// Pitching amplitude (degrees) about x, y, & z axes (RIGID_MOTION only)
	PitchingAmplX []float64
	// Pitching amplitude (degrees) about x, y, & z axes (RIGID_MOTION only)
	PitchingAmplY []float64
	// Pitching amplitude (degrees) about x, y, & z axes (RIGID_MOTION only)
	PitchingAmplZ []float64
	// Pitching phase offset (degrees) about x, y, & z axes (RIGID_MOTION only)
	PitchingPhaseX []float64
	// Pitching phase offset (degrees) about x, y, & z axes (RIGID_MOTION only)
	PitchingPhaseY []float64
	// Pitching phase offset (degrees) about x, y, & z axes (RIGID_MOTION only)
	PitchingPhaseZ []float64
	// Plunging angular freq. (rad/s) in x, y, & z directions (RIGID_MOTION only)
	PlungingOmegaX []float64
	// Plunging angular freq. (rad/s) in x, y, & z directions (RIGID_MOTION only)
	PlungingOmegaY []float64
	// Plunging angular freq. (rad/s) in x, y, & z directions (RIGID_MOTION only)
	PlungingOmegaZ []float64
	// Plunging amplitude (m) in x, y, & z directions (RIGID_MOTION only)
	PlungingAmplX []float64
	// Plunging amplitude (m) in x, y, & z directions (RIGID_MOTION only)
	PlungingAmplY []float64
	// Plunging amplitude (m) in x, y, & z directions (RIGID_MOTION only)
	PlungingAmplZ []float64
	// Value to move motion origins (1 or 0)
	MoveMotionOrigin []uint16
	//
	MotionFilename string
	// Uncoupled Aeroelastic Frequency Plunge.
	FreqPlungeAeroelastic float64
	// Uncoupled Aeroelastic Frequency Pitch.
	FreqPitchAeroelastic float64
	// Kind of grid adaptation
	KindAdapt enum.Adapt
	// Percentage of new elements (% of the original number of elements)
	NewElems float64
	// Scale factor for the dual volume
	DualvolPower float64
	// Use analytical definition for surfaces
	AnalyticalSurfdef enum.GeoAnalytic
	// Before each computation, implicitly smooth the nodal coordinates
	SmoothGeometry bool
	// Adapt the boundary elements
	AdaptBoundary bool
	// Divide rectangles into triangles
	DivideElements bool
	// Apply a wind gust
	WindGust bool
	// Type of gust
	GustType enum.GustType
	// Gust wavelenght (meters)
	GustWavelength float64
	// Number of gust periods
	GustPeriods float64
	// Gust amplitude (m/s)
	GustAmpl float64
	// Time at which to begin the gust (sec)
	GustBeginTime float64
	// Location at which the gust begins (meters)
	GustBeginLoc float64
	// Direction of the gust X or Y dir
	GustDir enum.GustDir
	// Evaluate equivalent area on the Near-Field
	EquivArea bool
	// Integration limits of the equivalent area ( xmin, xmax, Dist_NearField )
	EaIntLimit [3]float64
	// Specify chemical model for multi-species simulations
	GasModel enum.Gasmodel
	//
	GasComposition []float64
	// Ratio of density for two phase problems
	RatioDensity float64
	// Ratio of viscosity for two phase problems
	RatioViscosity float64
	// Location of the freesurface (y or z coordinate)
	FreesurfaceZero float64
	// Free surface depth surface (x or y coordinate)
	FreesurfaceDepth float64
	// Thickness of the interface in a free surface problem
	FreesurfaceThickness float64
	// Free surface damping coefficient
	FreesurfaceDampingCoeff float64
	// Free surface damping length (times the baseline wave)
	FreesurfaceDampingLength float64
	// Location of the free surface outlet surface (x or y coordinate)
	FreesurfaceOutlet float64
	// Kind of deformation
	DvKind []enum.Param
	// Marker of the surface to which we are going apply the shape deformation
	DvMarker []string
	// New value of the shape deformation
	DvValue []float64
	// Parameters of the shape deformation- FFD_CONTROL_POINT_2D ( FFDBox ID, i_Ind, j_Ind, x_Disp, y_Disp )- FFD_CAMBER_2D ( FFDBox ID, i_Ind )- FFD_THICKNESS_2D ( FFDBox ID, i_Ind )- HICKS_HENNE ( Lower Surface (0)/Upper Surface (1)/Only one Surface (2), x_Loc )- COSINE_BUMP ( Lower Surface (0)/Upper Surface (1)/Only one Surface (2), x_Loc, Thickness )- FOURIER ( Lower Surface (0)/Upper Surface (1)/Only one Surface (2), index, cos(0)/sin(1) )- NACA_4DIGITS ( 1st digit, 2nd digit, 3rd and 4th digit )- PARABOLIC ( Center, Thickness )- DISPLACEMENT ( x_Disp, y_Disp, z_Disp )- ROTATION ( x_Orig, y_Orig, z_Orig, x_End, y_End, z_End )- OBSTACLE ( Center, Bump size )- SPHERICAL ( ControlPoint_Index, Theta_Disp, R_Disp )- FFD_CONTROL_POINT ( FFDBox ID, i_Ind, j_Ind, k_Ind, x_Disp, y_Disp, z_Disp )- FFD_DIHEDRAL_ANGLE ( FFDBox ID, x_Orig, y_Orig, z_Orig, x_End, y_End, z_End )- FFD_TWIST_ANGLE ( FFDBox ID, x_Orig, y_Orig, z_Orig, x_End, y_End, z_End )- FFD_ROTATION ( FFDBox ID, x_Orig, y_Orig, z_Orig, x_End, y_End, z_End )- FFD_CONTROL_SURFACE ( FFDBox ID, x_Orig, y_Orig, z_Orig, x_End, y_End, z_End )- FFD_CAMBER ( FFDBox ID, i_Ind, j_Ind )- FFD_THICKNESS ( FFDBox ID, i_Ind, j_Ind )
	DvParam *su2types.DVParam
	// Hold the grid fixed in a region
	HoldGridFixed bool
	// Coordinates of the box where the grid will be deformed (Xmin, Ymin, Zmin, Xmax, Ymax, Zmax)
	HoldGridFixedCoord [6]float64
	// Visualize the deformation
	VisualizeDeformation bool
	// Print the residuals during mesh deformation to the console
	DeformConsoleOutput bool
	// Number of nonlinear deformation iterations (surface deformation increments)
	DeformNonlinearIter uint64
	// Number of smoothing iterations for FEA mesh deformation
	DeformLinearIter uint64
	// Factor to multiply smallest volume for deform tolerance (0.001 default)
	DeformTolFactor float64
	// Type of element stiffness imposed for FEA mesh deformation (INVERSE_VOLUME, WALL_DISTANCE, CONSTANT_STIFFNESS)
	DeformStiffnessType enum.DeformStiffness
	// Poisson's ratio for constant stiffness FEA method of grid deformation
	DeformElasticityModulus float64
	// Young's modulus and Poisson's ratio for constant stiffness FEA method of grid deformation
	DeformPoissonsRatio float64
	// MISSING ---
	CyclicPitch float64
	// MISSING ---
	CollectivePitch float64
	// Modulus of elasticity
	ElasticityModulus float64
	// Poisson ratio
	PoissonRatio float64
	// Material density
	MaterialDensity float64
	// Constant wave speed
	WaveSpeed float64
	// Thermal diffusivity constant
	ThermalDiffusivity float64
	// Node number for the CV to be visualized
	VisualizeCv int
	// Evaluate inverse design on the surface
	InvDesignCp bool
	// Evaluate inverse design on the surface
	InvDesignHeatflux bool
	// Write extra output
	ExtraOutput bool
	// Location of the turb model itself
	MlTurbModelFile string
	// what kind of input/output feature map is there
	MlTurbModelFeatureset string
	// Extra values for ML Turb model
	MlTurbModelExtra []string
	// Number of total iterations in the FFD point inversion
	FfdIterations uint16
	// Free surface damping coefficient
	FfdTolerance float64
	// Gradient method
	GradientMethod *su2types.Python
	// Geometrical Parameter
	GeoParam *su2types.Python
	// Setup for design variables
	DefinitionDv *su2types.Python
	// Current value of the design variables
	DvValueNew *su2types.Python
	// Previous value of the design variables
	DvValueOld *su2types.Python
	// Number of partitions of the mesh
	NumberPart *su2types.Python
	// Optimization objective function with optional scaling factor
	OptObjective *su2types.Python
	// Optimization constraint functions with optional scaling factor
	OptConstraint *su2types.Python
	// Finite different step for gradient estimation
	FinDiffStep *su2types.Python
	// Verbosity of the python scripts to Stdout
	Console *su2types.Python
	// Flag specifying if the mesh was decomposed
	Decomposed *su2types.Python
}
