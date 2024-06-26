// Code generated by running "go generate" in github.com/neovim/go-client/nvim. DO NOT EDIT.

package nvim

// EmbedOptions specifies options for starting an embedded instance of Nvim.
//
// Deprecated: Use ChildProcessOption instead.
type EmbedOptions struct {
	// Logf log function for rpc.WithLogf.
	Logf func(string, ...any)

	// Dir specifies the working directory of the command. The working
	// directory in the current process is used if Dir is "".
	Dir string

	// Path is the path of the command to run. If Path = "", then
	// StartEmbeddedNvim searches for "nvim" on $PATH.
	Path string

	// Args specifies the command line arguments. Do not include the program
	// name (the first argument) or the --embed option.
	Args []string

	// Env specifies the environment of the Nvim process. The current process
	// environment is used if Env is nil.
	Env []string
}

// NewEmbedded starts an embedded instance of Nvim using the specified options.
//
// The application must call Serve() to handle RPC requests and responses.
//
// Deprecated: Use NewChildProcess instead.
func NewEmbedded(options *EmbedOptions) (*Nvim, error) {
	if options == nil {
		options = &EmbedOptions{}
	}
	path := options.Path
	if path == "" {
		path = "nvim"
	}

	return NewChildProcess(
		ChildProcessArgs(append([]string{"--embed"}, options.Args...)...),
		ChildProcessCommand(path),
		ChildProcessEnv(options.Env),
		ChildProcessDir(options.Dir),
		ChildProcessServe(false))
}

// ExecuteLua executes a Lua block.
//
// Deprecated: Use ExecLua instead.
func (v *Nvim) ExecuteLua(code string, result any, args ...any) error {
	if args == nil {
		args = emptyArgs
	}
	return v.call("nvim_execute_lua", result, code, args)
}

// ExecuteLua executes a Lua block.
//
// Deprecated: Use ExecLua instead.
func (b *Batch) ExecuteLua(code string, result any, args ...any) {
	if args == nil {
		args = emptyArgs
	}
	b.call("nvim_execute_lua", result, code, args)
}

// BufferNumber gets a buffer's number.
//
// Deprecated: Use int(buffer) to get the buffer's number as an integer.
//
// See: [nvim_buf_get_number()]
//
// [nvim_buf_get_number()]: https://neovim.io/doc/user/api.html#nvim_buf_get_number()
func (v *Nvim) BufferNumber(buffer Buffer) (number int, err error) {
	err = v.call("nvim_buf_get_number", &number, buffer)
	return number, err
}

// BufferNumber gets a buffer's number.
//
// Deprecated: Use int(buffer) to get the buffer's number as an integer.
//
// See: [nvim_buf_get_number()]
//
// [nvim_buf_get_number()]: https://neovim.io/doc/user/api.html#nvim_buf_get_number()
func (b *Batch) BufferNumber(buffer Buffer, number *int) {
	b.call("nvim_buf_get_number", number, buffer)
}

// ClearBufferHighlight clears highlights from a given source group and a range
// of lines.
//
// To clear a source group in the entire buffer, pass in 1 and -1 to startLine
// and endLine respectively.
//
// The lineStart and lineEnd parameters specify the range of lines to clear.
// The end of range is exclusive. Specify -1 to clear to the end of the file.
//
// Deprecated: Use ClearBufferNamespace instead.
//
// See: [nvim_buf_clear_highlight()]
//
// [nvim_buf_clear_highlight()]: https://neovim.io/doc/user/api.html#nvim_buf_clear_highlight()
func (v *Nvim) ClearBufferHighlight(buffer Buffer, srcID int, startLine int, endLine int) error {
	return v.call("nvim_buf_clear_highlight", nil, buffer, srcID, startLine, endLine)
}

// ClearBufferHighlight clears highlights from a given source group and a range
// of lines.
//
// To clear a source group in the entire buffer, pass in 1 and -1 to startLine
// and endLine respectively.
//
// The lineStart and lineEnd parameters specify the range of lines to clear.
// The end of range is exclusive. Specify -1 to clear to the end of the file.
//
// Deprecated: Use ClearBufferNamespace instead.
//
// See: [nvim_buf_clear_highlight()]
//
// [nvim_buf_clear_highlight()]: https://neovim.io/doc/user/api.html#nvim_buf_clear_highlight()
func (b *Batch) ClearBufferHighlight(buffer Buffer, srcID int, startLine int, endLine int) {
	b.call("nvim_buf_clear_highlight", nil, buffer, srcID, startLine, endLine)
}

// SetBufferVirtualText set the virtual text (annotation) for a buffer line.
//
// By default (and currently the only option), the text will be placed after
// the buffer text.
//
// Virtual text will never cause reflow, rather virtual text will be truncated at the end of the screen line.
// The virtual text will begin one cell (|lcs-eol| or space) after the ordinary text.
//
// Namespaces are used to support batch deletion/updating of virtual text.
// To create a namespace, use CreateNamespace. Virtual text is cleared using ClearBufferNamespace.
//
// The same nsID can be used for both virtual text and highlights added by AddBufferHighlight,
// both can then be cleared with a single call to ClearBufferNamespace.
// If the virtual text never will be cleared by an API call, pass "nsID = -1".
//
// As a shorthand, "nsID = 0" can be used to create a new namespace for the
// virtual text, the allocated id is then returned.
//
// The opts arg is reserved for future use.
//
// Deprecated: Use SetBufferExtmark instead.
//
// See: [nvim_buf_set_virtual_text()]
//
// [nvim_buf_set_virtual_text()]: https://neovim.io/doc/user/api.html#nvim_buf_set_virtual_text()
func (v *Nvim) SetBufferVirtualText(buffer Buffer, nsID int, line int, chunks []TextChunk, opts map[string]any) (id int, err error) {
	err = v.call("nvim_buf_set_virtual_text", &id, buffer, nsID, line, chunks, opts)
	return id, err
}

// SetBufferVirtualText set the virtual text (annotation) for a buffer line.
//
// By default (and currently the only option), the text will be placed after
// the buffer text.
//
// Virtual text will never cause reflow, rather virtual text will be truncated at the end of the screen line.
// The virtual text will begin one cell (|lcs-eol| or space) after the ordinary text.
//
// Namespaces are used to support batch deletion/updating of virtual text.
// To create a namespace, use CreateNamespace. Virtual text is cleared using ClearBufferNamespace.
//
// The same nsID can be used for both virtual text and highlights added by AddBufferHighlight,
// both can then be cleared with a single call to ClearBufferNamespace.
// If the virtual text never will be cleared by an API call, pass "nsID = -1".
//
// As a shorthand, "nsID = 0" can be used to create a new namespace for the
// virtual text, the allocated id is then returned.
//
// The opts arg is reserved for future use.
//
// Deprecated: Use SetBufferExtmark instead.
//
// See: [nvim_buf_set_virtual_text()]
//
// [nvim_buf_set_virtual_text()]: https://neovim.io/doc/user/api.html#nvim_buf_set_virtual_text()
func (b *Batch) SetBufferVirtualText(buffer Buffer, nsID int, line int, chunks []TextChunk, opts map[string]any, id *int) {
	b.call("nvim_buf_set_virtual_text", id, buffer, nsID, line, chunks, opts)
}

// HLByID gets a highlight definition by name.
//
// hlID is the highlight id as returned by HLIDByName.
//
// rgb is the whether the export RGB colors.
//
// The returned highlight is the highlight definition.
//
// See: [nvim_get_hl_by_id()]
//
// [nvim_get_hl_by_id()]: https://neovim.io/doc/user/api.html#nvim_get_hl_by_id()
func (v *Nvim) HLByID(hlID int, rgb bool) (highlight *HLAttrs, err error) {
	var result HLAttrs
	err = v.call("nvim_get_hl_by_id", &result, hlID, rgb)
	return &result, err
}

// HLByID gets a highlight definition by name.
//
// hlID is the highlight id as returned by HLIDByName.
//
// rgb is the whether the export RGB colors.
//
// The returned highlight is the highlight definition.
//
// See: [nvim_get_hl_by_id()]
//
// [nvim_get_hl_by_id()]: https://neovim.io/doc/user/api.html#nvim_get_hl_by_id()
func (b *Batch) HLByID(hlID int, rgb bool, highlight *HLAttrs) {
	b.call("nvim_get_hl_by_id", highlight, hlID, rgb)
}

// HLByName gets a highlight definition by id.
//
// name is Highlight group name.
//
// rgb is whether the export RGB colors.
//
// The returned highlight is the highlight definition.
//
// See: [nvim_get_hl_by_name()]
//
// [nvim_get_hl_by_name()]: https://neovim.io/doc/user/api.html#nvim_get_hl_by_name()
func (v *Nvim) HLByName(name string, rgb bool) (highlight *HLAttrs, err error) {
	var result HLAttrs
	err = v.call("nvim_get_hl_by_name", &result, name, rgb)
	return &result, err
}

// HLByName gets a highlight definition by id.
//
// name is Highlight group name.
//
// rgb is whether the export RGB colors.
//
// The returned highlight is the highlight definition.
//
// See: [nvim_get_hl_by_name()]
//
// [nvim_get_hl_by_name()]: https://neovim.io/doc/user/api.html#nvim_get_hl_by_name()
func (b *Batch) HLByName(name string, rgb bool, highlight *HLAttrs) {
	b.call("nvim_get_hl_by_name", highlight, name, rgb)
}

// CommandOutput executes a single ex command and returns the output.
//
// Deprecated: Use Exec instead.
//
// See: [nvim_command_output()]
//
// [nvim_command_output()]: https://neovim.io/doc/user/api.html#nvim_command_output()
func (v *Nvim) CommandOutput(cmd string) (out string, err error) {
	err = v.call("nvim_command_output", &out, cmd)
	return out, err
}

// CommandOutput executes a single ex command and returns the output.
//
// Deprecated: Use Exec instead.
//
// See: [nvim_command_output()]
//
// [nvim_command_output()]: https://neovim.io/doc/user/api.html#nvim_command_output()
func (b *Batch) CommandOutput(cmd string, out *string) {
	b.call("nvim_command_output", out, cmd)
}
