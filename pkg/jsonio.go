package packlit

import (
	"encoding/json"
	"fmt"
)

// StreamDescriptor represents a single stream descriptor in JSON.
// It closely maps to a descriptor built by WithInput/WithStream/WithOutput and other options.
type StreamDescriptor struct {
	Input           string            `json:"input"`
	Stream          string            `json:"stream"`
	Output          string            `json:"output"`
	InitSegment     string            `json:"init_segment,omitempty"`
	SegmentTemplate string            `json:"segment_template,omitempty"`
	Extra           map[string]string `json:"extra,omitempty"`
}

// PackagerOptions is the top-level JSON payload accepted by parser API.
type PackagerOptions struct {
	Binary     string             `json:"binary,omitempty"`
	Streams    []StreamDescriptor `json:"streams"`
	Packager   *PackagerFlags     `json:"packager,omitempty"`
	Muxer      *MuxerFlags        `json:"muxer,omitempty"`
	MPD        *MpdFlags          `json:"mpd,omitempty"`
	Manifest   *ManifestFlags     `json:"manifest,omitempty"`
	HTTP       *HttpFlags         `json:"http,omitempty"`
	HLS        *HlsFlags          `json:"hls,omitempty"`
	Encryption *EncryptionFlags   `json:"encryption,omitempty"`
}

// Flag configuration structs
type PackagerFlags struct {
	Licenses             bool   `json:"licenses,omitempty"`
	Quiet                bool   `json:"quiet,omitempty"`
	UseFakeClockForMuxer bool   `json:"use_fake_clock_for_muxer,omitempty"`
	TestPackagerVersion  bool   `json:"test_packager_version,omitempty"`
	SingleThreaded       bool   `json:"single_threaded,omitempty"`
	VModule              string `json:"vmodule,omitempty"`
	Version              bool   `json:"version,omitempty"`
}

type MuxerFlags struct {
	ClearLead                        float32 `json:"clear_lead,omitempty"`
	SegmentDuration                  float32 `json:"segment_duration,omitempty"`
	SegmentSapAligned                bool    `json:"segment_sap_aligned,omitempty"`
	FragmentDuration                 float32 `json:"fragment_duration,omitempty"`
	FragmentSapAligned               bool    `json:"fragment_sap_aligned,omitempty"`
	GenerateSidxInMediaSegments      bool    `json:"generate_sidx_in_media_segments,omitempty"`
	TempDir                          string  `json:"temp_dir,omitempty"`
	Mp4IncludePsshInStream           bool    `json:"mp4_include_pssh_in_stream,omitempty"`
	TransportStreamTimestampOffsetMs int     `json:"transport_stream_timestamp_offset_ms,omitempty"`
	DefaultTextZeroBiasMs            int     `json:"default_text_zero_bias_ms,omitempty"`
	StartSegmentNumber               int     `json:"start_segment_number,omitempty"`
	UseDoviSupplementalCodecs        bool    `json:"use_dovi_supplemental_codecs,omitempty"`
}

type MpdFlags struct {
	OutputMediaInfo                 bool              `json:"output_media_info,omitempty"`
	MpdOutput                       string            `json:"mpd_output,omitempty"`
	BaseUrls                        []string          `json:"base_urls,omitempty"`
	GenerateStaticLiveMpd           bool              `json:"generate_static_live_mpd,omitempty"`
	AllowApproximateSegmentTimeline bool              `json:"allow_approximate_segment_timeline,omitempty"`
	DumpStreamInfo                  bool              `json:"dump_stream_info,omitempty"`
	MinBufferTime                   float32           `json:"min_buffer_time,omitempty"`
	MinimumUpdatePeriod             float32           `json:"minimum_update_period,omitempty"`
	SuggestedPresentationDelay      float32           `json:"suggested_presentation_delay,omitempty"`
	UtcTimings                      map[string]string `json:"utc_timings,omitempty"`
	GenerateDashIfIopCompliantMpd   bool              `json:"generate_dash_if_iop_compliant_mpd,omitempty"`
	AllowCodecSwitching             bool              `json:"allow_codec_switching,omitempty"`
	IncludeMsprProForPlayReady      bool              `json:"include_mspr_pro_for_playready,omitempty"`
	DashForceSegmentList            bool              `json:"dash_force_segment_list,omitempty"`
	LowLatencyDashMode              bool              `json:"low_latency_dash_mode,omitempty"`
}

type ManifestFlags struct {
	TimeShiftBufferDepth               float32 `json:"time_shift_buffer_depth,omitempty"`
	PreservedSegmentsOutsideLiveWindow int     `json:"preserved_segments_outside_live_window,omitempty"`
	DefaultLanguage                    string  `json:"default_language,omitempty"`
	DefaultTextLanguage                string  `json:"default_text_language,omitempty"`
	ForceClIndex                       bool    `json:"force_cl_index,omitempty"`
}

type HttpFlags struct {
	UserAgent                    string `json:"user_agent,omitempty"`
	CaFile                       string `json:"ca_file,omitempty"`
	ClientCertFile               string `json:"client_cert_file,omitempty"`
	ClientCertPrivateKeyFile     string `json:"client_cert_private_key_file,omitempty"`
	ClientCertPrivateKeyPassword string `json:"client_cert_private_key_password,omitempty"`
	DisablePeerVerification      bool   `json:"disable_peer_verification,omitempty"`
	IgnoreHttpOutputFailures     bool   `json:"ignore_http_output_failures,omitempty"`
	IoCacheSize                  uint   `json:"io_cache_size,omitempty"`
}

type HlsFlags struct {
	MasterPlaylistOutput string  `json:"hls_master_playlist_output,omitempty"`
	BaseUrl              string  `json:"hls_base_url,omitempty"`
	KeyUri               string  `json:"hls_key_uri,omitempty"`
	PlaylistType         string  `json:"hls_playlist_type,omitempty"`
	MediaSequenceNumber  int     `json:"hls_media_sequence_number,omitempty"`
	StartTimeOffset      float32 `json:"hls_start_time_offset,omitempty"`
	CreateSessionKeys    bool    `json:"create_session_keys,omitempty"`
}

type EncryptionFlags struct {
	CryptByteBlock           int             `json:"crypt_byte_block,omitempty"`
	ProtectionScheme         string          `json:"protection_scheme,omitempty"`
	ProtectionSystems        []string        `json:"protection_systems,omitempty"`
	SkipByteBlock            int             `json:"skip_byte_block,omitempty"`
	Vp9SubsampleEncryption   bool            `json:"vp9_subsample_encryption,omitempty"`
	PlayreadyExtraHeaderData string          `json:"playready_extra_header_data,omitempty"`
	RawKey                   *RawKeyFlags    `json:"raw_key,omitempty"`
	Widevine                 *WideVineFlags  `json:"widevine,omitempty"`
	Playready                *PlayReadyFlags `json:"playready,omitempty"`
}

type RawKeyInfo struct {
	Label string `json:"label"`
	KeyId string `json:"key_id"`
	Key   string `json:"key"`
}

func (r *RawKeyInfo) String() string {
	return fmt.Sprintf("label=%s:key_id=%s:key=%s", r.Label, r.KeyId, r.Key)
}

type RawKeyFlags struct {
	EnableEncryption bool         `json:"enable_raw_key_encryption,omitempty"`
	EnableDecryption bool         `json:"enable_raw_key_decryption,omitempty"`
	Keys             []RawKeyInfo `json:"keys,omitempty"`
	Iv               string       `json:"iv,omitempty"`
	Pssh             string       `json:"pssh,omitempty"`
}

type WideVineFlags struct {
	EnableEncryption     bool   `json:"enable_widevine_encryption,omitempty"`
	EnableDecryption     bool   `json:"enable_widevine_decryption,omitempty"`
	KeyServerUrl         string `json:"key_server_url,omitempty"`
	ContentId            string `json:"content_id,omitempty"`
	Policy               string `json:"policy,omitempty"`
	MaxSdPixels          int    `json:"max_sd_pixels,omitempty"`
	MaxHdPixels          int    `json:"max_hd_pixels,omitempty"`
	MaxUhd1Pixels        int    `json:"max_uhd1_pixels,omitempty"`
	Signer               string `json:"signer,omitempty"`
	AesSigningKey        string `json:"aes_signing_key,omitempty"`
	AesSigningIv         string `json:"aes_signing_iv,omitempty"`
	RsaSigningKeyPath    string `json:"rsa_signing_key_path,omitempty"`
	CryptoPeriodDuration int    `json:"crypto_period_duration,omitempty"`
	GroupId              string `json:"group_id,omitempty"`
	EnableEntitlement    bool   `json:"enable_entitlement_license,omitempty"`
}

type PlayReadyFlags struct {
	EnableEncryption  bool   `json:"enable_playready_encryption,omitempty"`
	ServerUrl         string `json:"playready_server_url,omitempty"`
	ProgramIdentifier string `json:"program_identifier,omitempty"`
}

// FlagProcessor interface for processing different flag types
type FlagProcessor interface {
	ProcessFlags() ([]ShakaFlagFn, error)
}

var _ FlagProcessor = (*PackagerFlagsProcessor)(nil)

// PackagerFlagsProcessor processes packager flags
type PackagerFlagsProcessor struct {
	flags *PackagerFlags
}

func NewPackagerFlagsProcessor(flags *PackagerFlags) *PackagerFlagsProcessor {
	return &PackagerFlagsProcessor{flags: flags}
}

func (p *PackagerFlagsProcessor) ProcessFlags() ([]ShakaFlagFn, error) {
	if p.flags == nil {
		return nil, nil
	}

	var flagFns []ShakaFlagFn

	if p.flags.Licenses {
		flagFns = append(flagFns, WithLicensesFlag())
	}
	if p.flags.Quiet {
		flagFns = append(flagFns, WithQuietFlag())
	}
	if p.flags.UseFakeClockForMuxer {
		flagFns = append(flagFns, WithUseFakeClockForMuxerFlag())
	}
	if p.flags.TestPackagerVersion {
		flagFns = append(flagFns, WithTestPackagerVersionFlag())
	}
	if p.flags.SingleThreaded {
		flagFns = append(flagFns, WithSingleThreadedFlag())
	}
	if p.flags.VModule != "" {
		flagFns = append(flagFns, WithVModuleFlag(p.flags.VModule))
	}
	if p.flags.Version {
		flagFns = append(flagFns, WithVersionFlag())
	}

	return flagFns, nil
}

// MuxerFlagsProcessor processes muxer flags
type MuxerFlagsProcessor struct {
	flags *MuxerFlags
}

var _ FlagProcessor = (*MuxerFlagsProcessor)(nil)

func NewMuxerFlagsProcessor(flags *MuxerFlags) *MuxerFlagsProcessor {
	return &MuxerFlagsProcessor{flags: flags}
}

func (m *MuxerFlagsProcessor) ProcessFlags() ([]ShakaFlagFn, error) {
	if m.flags == nil {
		return nil, nil
	}

	var flagFns []ShakaFlagFn

	if m.flags.ClearLead != 0 {
		flagFns = append(flagFns, WithClearLeadFlag(m.flags.ClearLead))
	}
	if m.flags.SegmentDuration != 0 {
		flagFns = append(flagFns, WithSegmentDurationFlag(m.flags.SegmentDuration))
	}
	if m.flags.SegmentSapAligned {
		flagFns = append(flagFns, WithSegmentSapAlignedFlag())
	}
	if m.flags.FragmentDuration != 0 {
		flagFns = append(flagFns, WithFragmentDurationFlag(m.flags.FragmentDuration))
	}
	if m.flags.FragmentSapAligned {
		flagFns = append(flagFns, WithFragmentSapAlignedFlag())
	}
	if m.flags.GenerateSidxInMediaSegments {
		flagFns = append(flagFns, WithGenerateSidxInMediaSegmentsFlag())
	}
	if m.flags.TempDir != "" {
		flagFns = append(flagFns, WithTempDirFlag(m.flags.TempDir))
	}
	if m.flags.Mp4IncludePsshInStream {
		flagFns = append(flagFns, WithMp4IncludePsshInStreamFlag())
	}
	if m.flags.TransportStreamTimestampOffsetMs != 0 {
		flagFns = append(flagFns, WithTransportStreamTimestampOffsetMsFlag(m.flags.TransportStreamTimestampOffsetMs))
	}
	if m.flags.DefaultTextZeroBiasMs != 0 {
		flagFns = append(flagFns, WithDefaultTextZeroBiasMsFlag(m.flags.DefaultTextZeroBiasMs))
	}
	if m.flags.StartSegmentNumber != 0 {
		flagFns = append(flagFns, WithStartSegmentNumberFlag(m.flags.StartSegmentNumber))
	}
	if m.flags.UseDoviSupplementalCodecs {
		flagFns = append(flagFns, WithUseDoviSupplementalCodecsFlag())
	}

	return flagFns, nil
}

// MPDFlagsProcessor processes MPD flags
type MPDFlagsProcessor struct {
	flags *MpdFlags
}

var _ FlagProcessor = (*MPDFlagsProcessor)(nil)

func NewMPDFlagsProcessor(flags *MpdFlags) *MPDFlagsProcessor {
	return &MPDFlagsProcessor{flags: flags}
}

func (m *MPDFlagsProcessor) ProcessFlags() ([]ShakaFlagFn, error) {
	if m.flags == nil {
		return nil, nil
	}

	var flagFns []ShakaFlagFn

	if m.flags.OutputMediaInfo {
		flagFns = append(flagFns, WithOutputMediaInfoFlag())
	}
	if m.flags.MpdOutput != "" {
		flagFns = append(flagFns, WithMpdOutput(m.flags.MpdOutput))
	}
	if len(m.flags.BaseUrls) > 0 {
		flagFns = append(flagFns, WithBaseUrls(m.flags.BaseUrls))
	}
	if m.flags.GenerateStaticLiveMpd {
		flagFns = append(flagFns, WithStaticLiveMpd())
	}
	if m.flags.AllowApproximateSegmentTimeline {
		flagFns = append(flagFns, WithAllowApproximateSegmentTimelineFlag())
	}
	if m.flags.DumpStreamInfo {
		flagFns = append(flagFns, WithDumpStream())
	}
	if m.flags.MinBufferTime != 0 {
		flagFns = append(flagFns, WithMinBufferTimeFlag(m.flags.MinBufferTime))
	}
	if m.flags.MinimumUpdatePeriod != 0 {
		flagFns = append(flagFns, WithMinimumUpdatePeriodFlag(m.flags.MinimumUpdatePeriod))
	}
	if m.flags.SuggestedPresentationDelay != 0 {
		flagFns = append(flagFns, WithSuggestedPresentationDelayFlag(m.flags.SuggestedPresentationDelay))
	}
	if len(m.flags.UtcTimings) > 0 {
		// Convert map to comma-separated scheme_id_uri=value pairs
		timingPairs := make([]string, 0, len(m.flags.UtcTimings))
		for schemeIdUri, value := range m.flags.UtcTimings {
			if schemeIdUri == "" {
				return nil, fmt.Errorf("empty scheme_id_uri found in utc_timings")
			}
			timingPairs = append(timingPairs, fmt.Sprintf("%s=%s", schemeIdUri, value))
		}
		flagFns = append(flagFns, WithUtCTimingsFlag(timingPairs))
	}
	if m.flags.GenerateDashIfIopCompliantMpd {
		flagFns = append(flagFns, WithGenerateDashIfIopCompliantMpdFlag())
	}
	if m.flags.AllowCodecSwitching {
		flagFns = append(flagFns, WithAllowCodecSwitchingFlag())
	}
	if m.flags.IncludeMsprProForPlayReady {
		flagFns = append(flagFns, WithIncludeMsprProForPlayReadyFlag())
	}
	if m.flags.DashForceSegmentList {
		flagFns = append(flagFns, WithDashForceSegmentListFlag())
	}
	if m.flags.LowLatencyDashMode {
		flagFns = append(flagFns, WithLowLatencyDashModeFlag(true))
	}

	return flagFns, nil
}

// ManifestFlagsProcessor processes manifest flags
type ManifestFlagsProcessor struct {
	flags *ManifestFlags
}

var _ FlagProcessor = (*ManifestFlagsProcessor)(nil)

func NewManifestFlagsProcessor(flags *ManifestFlags) *ManifestFlagsProcessor {
	return &ManifestFlagsProcessor{flags: flags}
}

func (m *ManifestFlagsProcessor) ProcessFlags() ([]ShakaFlagFn, error) {
	if m.flags == nil {
		return nil, nil
	}

	var flagFns []ShakaFlagFn

	if m.flags.TimeShiftBufferDepth != 0 {
		flagFns = append(flagFns, WithTimeShiftBufferDepthFlag(m.flags.TimeShiftBufferDepth))
	}
	if m.flags.PreservedSegmentsOutsideLiveWindow != 0 {
		flagFns = append(flagFns, WithPreservedSegmentsOutsideLiveWindowFlag(m.flags.PreservedSegmentsOutsideLiveWindow))
	}
	if m.flags.DefaultLanguage != "" {
		flagFns = append(flagFns, WithDefaultLanguageFlag(m.flags.DefaultLanguage))
	}
	if m.flags.DefaultTextLanguage != "" {
		flagFns = append(flagFns, WithDefaultTextLanguageFlag(m.flags.DefaultTextLanguage))
	}
	if m.flags.ForceClIndex {
		flagFns = append(flagFns, WithForceCLIndexFlag())
	}

	return flagFns, nil
}

// HTTPFlagsProcessor processes HTTP flags
type HTTPFlagsProcessor struct {
	flags *HttpFlags
}

var _ FlagProcessor = (*HTTPFlagsProcessor)(nil)

func NewHTTPFlagsProcessor(flags *HttpFlags) *HTTPFlagsProcessor {
	return &HTTPFlagsProcessor{flags: flags}
}

func (h *HTTPFlagsProcessor) ProcessFlags() ([]ShakaFlagFn, error) {
	if h.flags == nil {
		return nil, nil
	}

	var flagFns []ShakaFlagFn

	if h.flags.UserAgent != "" {
		flagFns = append(flagFns, WithUserAgentFlag(h.flags.UserAgent))
	}
	if h.flags.CaFile != "" {
		flagFns = append(flagFns, WithCaFileFlag(h.flags.CaFile))
	}
	if h.flags.ClientCertFile != "" {
		flagFns = append(flagFns, WithClientCertFileFlag(h.flags.ClientCertFile))
	}
	if h.flags.ClientCertPrivateKeyFile != "" {
		flagFns = append(flagFns, WithClientCertPrivateKeyFileFlag(h.flags.ClientCertPrivateKeyFile))
	}
	if h.flags.ClientCertPrivateKeyPassword != "" {
		flagFns = append(flagFns, WithClientCertPrivateKeyPasswordFlag(h.flags.ClientCertPrivateKeyPassword))
	}
	if h.flags.DisablePeerVerification {
		flagFns = append(flagFns, WithDisablePeerVerificationFlag())
	}
	if h.flags.IgnoreHttpOutputFailures {
		flagFns = append(flagFns, WithIgnoreHttpOutputFailuresFlag())
	}
	if h.flags.IoCacheSize != 0 {
		flagFns = append(flagFns, WithIoCacheSizeFlag(h.flags.IoCacheSize))
	}

	return flagFns, nil
}

// HLSFlagsProcessor processes HLS flags
type HLSFlagsProcessor struct {
	flags *HlsFlags
}

var _ FlagProcessor = (*HLSFlagsProcessor)(nil)

func NewHLSFlagsProcessor(flags *HlsFlags) *HLSFlagsProcessor {
	return &HLSFlagsProcessor{flags: flags}
}

func (h *HLSFlagsProcessor) ProcessFlags() ([]ShakaFlagFn, error) {
	if h.flags == nil {
		return nil, nil
	}

	var flagFns []ShakaFlagFn

	if h.flags.MasterPlaylistOutput != "" {
		flagFns = append(flagFns, WithHLSMasterPlaylistOutputFlag(h.flags.MasterPlaylistOutput))
	}
	if h.flags.BaseUrl != "" {
		flagFns = append(flagFns, WithHLSBaseURLFlag(h.flags.BaseUrl))
	}
	if h.flags.KeyUri != "" {
		flagFns = append(flagFns, WithHLSKeyURIFlag(h.flags.KeyUri))
	}
	if h.flags.PlaylistType != "" {
		flagFns = append(flagFns, WithHLSPlaylistTypeFlag(h.flags.PlaylistType))
	}
	if h.flags.MediaSequenceNumber != 0 {
		flagFns = append(flagFns, WithHLSMediaSequenceNumberFlag(h.flags.MediaSequenceNumber))
	}
	if h.flags.StartTimeOffset != 0 {
		flagFns = append(flagFns, WithHLSStartTimeOffsetFlag(h.flags.StartTimeOffset))
	}
	if h.flags.CreateSessionKeys {
		flagFns = append(flagFns, WithHLSCreateSessionKeysFlag())
	}

	return flagFns, nil
}

// EncryptionFlagsProcessor processes crypto flags
type EncryptionFlagsProcessor struct {
	flags *EncryptionFlags
}

var _ FlagProcessor = (*EncryptionFlagsProcessor)(nil)

func NewEncryptionFlagsProcessor(flags *EncryptionFlags) *EncryptionFlagsProcessor {
	return &EncryptionFlagsProcessor{flags: flags}
}

func (c *EncryptionFlagsProcessor) ProcessFlags() ([]ShakaFlagFn, error) {
	if c.flags == nil {
		return nil, nil
	}

	var flagFns []ShakaFlagFn

	if c.flags.CryptByteBlock != 0 {
		flagFns = append(flagFns, WithCryptByteBlockFlag(c.flags.CryptByteBlock))
	}
	if c.flags.ProtectionScheme != "" {
		flagFns = append(flagFns, WithProtectionSchemeFlag(c.flags.ProtectionScheme))
	}
	if len(c.flags.ProtectionSystems) > 0 {
		flagFns = append(flagFns, WithProtectionSystemsFlag(c.flags.ProtectionSystems))
	}
	if c.flags.SkipByteBlock != 0 {
		flagFns = append(flagFns, WithSkipByteBlockFlag(c.flags.SkipByteBlock))
	}
	if c.flags.Vp9SubsampleEncryption {
		flagFns = append(flagFns, WithVP9SubsampleEncryptionFlag())
	}
	if c.flags.PlayreadyExtraHeaderData != "" {
		flagFns = append(flagFns, WithPlayReadyExtraHeaderDataFlag(c.flags.PlayreadyExtraHeaderData))
	}

	return flagFns, nil
}

// RawKeyFlagsProcessor processes raw key flags
type RawKeyFlagsProcessor struct {
	flags *RawKeyFlags
}

var _ FlagProcessor = (*RawKeyFlagsProcessor)(nil)

func NewRawKeyFlagsProcessor(flags *RawKeyFlags) *RawKeyFlagsProcessor {
	return &RawKeyFlagsProcessor{flags: flags}
}

func (r *RawKeyFlagsProcessor) ProcessFlags() ([]ShakaFlagFn, error) {
	if r.flags == nil {
		return nil, nil
	}

	var flagFns []ShakaFlagFn

	if r.flags.EnableEncryption {
		flagFns = append(flagFns, WithEnableRawKeyEncryptionFlag())
	}
	if r.flags.EnableDecryption {
		flagFns = append(flagFns, WithEnableRawKeyDecryptionFlag())
	}
	if len(r.flags.Keys) > 0 {
		// Validate and convert key info structs to formatted strings
		keyStrings := make([]string, 0, len(r.flags.Keys))
		for _, keyInfo := range r.flags.Keys {
			keyStrings = append(keyStrings, keyInfo.String())
		}
		flagFns = append(flagFns, WithKeysFlag(keyStrings))
	}
	if r.flags.Iv != "" {
		flagFns = append(flagFns, WithIvFlag(r.flags.Iv))
	}
	if r.flags.Pssh != "" {
		flagFns = append(flagFns, WithPsshFlag(r.flags.Pssh))
	}

	return flagFns, nil
}

// WideVineFlagsProcessor processes Widevine flags
type WideVineFlagsProcessor struct {
	flags *WideVineFlags
}

var _ FlagProcessor = (*WideVineFlagsProcessor)(nil)

func NewWideVineFlagsProcessor(flags *WideVineFlags) *WideVineFlagsProcessor {
	return &WideVineFlagsProcessor{flags: flags}
}

func (w *WideVineFlagsProcessor) ProcessFlags() ([]ShakaFlagFn, error) {
	if w.flags == nil {
		return nil, nil
	}

	var flagFns []ShakaFlagFn

	if w.flags.EnableEncryption {
		flagFns = append(flagFns, WithEnableWidevineEncryptionFlag())
	}
	if w.flags.EnableDecryption {
		flagFns = append(flagFns, WithEnableWidevineDecryptionFlag())
	}
	if w.flags.KeyServerUrl != "" {
		flagFns = append(flagFns, WithKeyServerURLFlag(w.flags.KeyServerUrl))
	}
	if w.flags.ContentId != "" {
		flagFns = append(flagFns, WithContentIDFlag(w.flags.ContentId))
	}
	if w.flags.Policy != "" {
		flagFns = append(flagFns, WithPolicyFlag(w.flags.Policy))
	}
	if w.flags.MaxSdPixels != 0 {
		flagFns = append(flagFns, WithMaxSDPixelsFlag(w.flags.MaxSdPixels))
	}
	if w.flags.MaxHdPixels != 0 {
		flagFns = append(flagFns, WithMaxHDPixelsFlag(w.flags.MaxHdPixels))
	}
	if w.flags.MaxUhd1Pixels != 0 {
		flagFns = append(flagFns, WithMaxUHD1PixelsFlag(w.flags.MaxUhd1Pixels))
	}
	if w.flags.Signer != "" {
		flagFns = append(flagFns, WithSignerFlag(w.flags.Signer))
	}
	if w.flags.AesSigningKey != "" {
		flagFns = append(flagFns, WithAESSigningKeyFlag(w.flags.AesSigningKey))
	}
	if w.flags.AesSigningIv != "" {
		flagFns = append(flagFns, WithAESSigningIVFlag(w.flags.AesSigningIv))
	}
	if w.flags.RsaSigningKeyPath != "" {
		flagFns = append(flagFns, WithRSASigningKeyPathFlag(w.flags.RsaSigningKeyPath))
	}
	if w.flags.CryptoPeriodDuration != 0 {
		flagFns = append(flagFns, WithCryptoPeriodDurationFlag(w.flags.CryptoPeriodDuration))
	}
	if w.flags.GroupId != "" {
		flagFns = append(flagFns, WithGroupIDFlag(w.flags.GroupId))
	}
	if w.flags.EnableEntitlement {
		flagFns = append(flagFns, WithEnableEntitlementLicenseFlag())
	}

	return flagFns, nil
}

// PlayReadyFlagsProcessor processes PlayReady flags
type PlayReadyFlagsProcessor struct {
	flags *PlayReadyFlags
}

var _ FlagProcessor = (*PlayReadyFlagsProcessor)(nil)

func NewPlayReadyFlagsProcessor(flags *PlayReadyFlags) *PlayReadyFlagsProcessor {
	return &PlayReadyFlagsProcessor{flags: flags}
}

func (p *PlayReadyFlagsProcessor) ProcessFlags() ([]ShakaFlagFn, error) {
	if p.flags == nil {
		return nil, nil
	}

	var flagFns []ShakaFlagFn

	if p.flags.EnableEncryption {
		flagFns = append(flagFns, WithPlayReadyEncryptionFlag())
	}
	if p.flags.ServerUrl != "" {
		flagFns = append(flagFns, WithPlayReadyServerURLFlag(p.flags.ServerUrl))
	}
	if p.flags.ProgramIdentifier != "" {
		flagFns = append(flagFns, WithProgramIdentifierFlag(p.flags.ProgramIdentifier))
	}

	return flagFns, nil
}

// PackagerJsonBuilder orchestrates the building process
type PackagerJsonBuilder struct {
	options *PackagerOptions
}

func NewPackagerJsonBuilder(options *PackagerOptions) *PackagerJsonBuilder {
	return &PackagerJsonBuilder{options: options}
}

// Build builds a ShakaPackager from PackagerOptions with comprehensive error handling
func (pb *PackagerJsonBuilder) Build() (*ShakaPackager, error) {
	builder := NewBuilder(pb.options.Binary)

	// Process all flag types using processors
	processors := []FlagProcessor{
		NewPackagerFlagsProcessor(pb.options.Packager),
		NewMuxerFlagsProcessor(pb.options.Muxer),
		NewMPDFlagsProcessor(pb.options.MPD),
		NewManifestFlagsProcessor(pb.options.Manifest),
		NewHTTPFlagsProcessor(pb.options.HTTP),
		NewHLSFlagsProcessor(pb.options.HLS),
		NewEncryptionFlagsProcessor(pb.options.Encryption),
	}

	// Add encryption-related processors if encryption options are provided
	if pb.options.Encryption != nil {
		processors = append(processors, NewRawKeyFlagsProcessor(pb.options.Encryption.RawKey))
		processors = append(processors, NewWideVineFlagsProcessor(pb.options.Encryption.Widevine))
		processors = append(processors, NewPlayReadyFlagsProcessor(pb.options.Encryption.Playready))
	}

	var allFlags []ShakaFlagFn
	for i, processor := range processors {
		flags, err := processor.ProcessFlags()
		if err != nil {
			return nil, fmt.Errorf("processor %d flag processing failed: %w", i, err)
		}

		allFlags = append(allFlags, flags...)
	}

	if len(allFlags) > 0 {
		builder.WithFlag(NewFlags(allFlags...))
	}

	// Build streams
	for _, s := range pb.options.Streams {
		streamBuilder := NewStreamBuilder()
		if s.Input != "" {
			streamBuilder = streamBuilder.WithInput(s.Input)
		}
		if s.Stream != "" {
			streamBuilder = streamBuilder.WithStream(StreamType(s.Stream))
		}
		if s.Output != "" {
			streamBuilder = streamBuilder.WithOutput(s.Output)
		}
		if s.InitSegment != "" {
			streamBuilder = streamBuilder.AddOption(WithInitSegment(s.InitSegment))
		}
		if s.SegmentTemplate != "" {
			streamBuilder = streamBuilder.AddOption(WithSegmentTemplate(s.SegmentTemplate))
		}
		// Extra arbitrary key=value pairs for stream descriptors
		if len(s.Extra) > 0 {
			for k, v := range s.Extra {
				key := k
				val := v
				streamBuilder = streamBuilder.AddOption(func(so *StreamOptions) {
					so.Add(GenericKeyValueOption{Key: key, Value: val})
				})
			}
		}

		builder.WithStream(streamBuilder.Build())
	}

	return builder.Build(), nil
}

// UnmarshalFlagsFromJSON is a helper to decode JSON payloads.
func UnmarshalFlagsFromJSON(data []byte) (PackagerOptions, error) {
	var req PackagerOptions
	if err := json.Unmarshal(data, &req); err != nil {
		return PackagerOptions{}, err
	}

	return req, nil
}

// BuildFromJSON is the main entry point with improved error handling
func BuildFromJSON(req PackagerOptions) (*ShakaPackager, error) {
	builder := NewPackagerJsonBuilder(&req)
	return builder.Build()
}
