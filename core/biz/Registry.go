package biz

func (r *Biz) RegisterBuiltinCommand(cmd BuiltInCommand) {
	r.BuiltInCommand[cmd.GetCaller()] = cmd
}

func (r *Biz) RegisterMagicCommand(cmd magicCommand) {
	r.MagicCommand[cmd.GetCaller()] = cmd
}

func (r *Biz) RegisterLuaPlugin(cmd luaPlugin) {
	r.LuaPlugin[cmd.GetCaller()] = cmd
}

func (r *Biz) GetBuiltinCommand(caller string) (BuiltInCommand, bool) {
	cmd, ok := r.BuiltInCommand[caller]
	return cmd, ok
}

func (r *Biz) GetMagicCommand(caller string) (BuiltInCommand, bool) {
	cmd, ok := r.MagicCommand[caller]
	return cmd, ok
}

func (r *Biz) GetLuaPlugin(caller string) (BuiltInCommand, bool) {
	cmd, ok := r.LuaPlugin[caller]
	return cmd, ok
}
