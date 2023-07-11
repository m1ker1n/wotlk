package core

import (
	"time"

	"github.com/wowsims/wotlk/sim/core/proto"
)

type APLValueCurrentRuneCount struct {
	defaultAPLValueImpl
	unit        *Unit
	ignoreDeath bool
	runeType    proto.APLValueRuneType
}

func (rot *APLRotation) newValueCurrentRuneCount(config *proto.APLValueCurrentRuneCount) APLValue {
	unit := rot.unit
	if !unit.HasRunicPowerBar() {
		return nil
	}
	return &APLValueCurrentRuneCount{
		unit:        unit,
		ignoreDeath: config.IgnoreDeath,
		runeType:    config.RuneType,
	}
}
func (value *APLValueCurrentRuneCount) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeInt
}
func (value *APLValueCurrentRuneCount) GetInt(sim *Simulation) int32 {
	switch value.runeType {
	case proto.APLValueRuneType_RuneBlood:
		if value.ignoreDeath {
			return int32(value.unit.NormalCurrentBloodRunes())
		} else {
			return int32(value.unit.CurrentBloodRunes())
		}
	case proto.APLValueRuneType_RuneFrost:
		if value.ignoreDeath {
			return int32(value.unit.NormalCurrentFrostRunes())
		} else {
			return int32(value.unit.CurrentFrostRunes())
		}
	case proto.APLValueRuneType_RuneUnholy:
		if value.ignoreDeath {
			return int32(value.unit.NormalCurrentUnholyRunes())
		} else {
			return int32(value.unit.CurrentUnholyRunes())
		}
	case proto.APLValueRuneType_RuneDeath:
		return int32(value.unit.CurrentDeathRunes())
	}
	return 0
}

type APLValueCurrentRuneActive struct {
	defaultAPLValueImpl
	unit     *Unit
	runeSlot int8
}

func (rot *APLRotation) newValueCurrentRuneActive(config *proto.APLValueCurrentRuneActive) APLValue {
	unit := rot.unit
	if !unit.HasRunicPowerBar() {
		return nil
	}
	return &APLValueCurrentRuneActive{
		unit:     unit,
		runeSlot: int8(config.RuneSlot),
	}
}
func (value *APLValueCurrentRuneActive) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeBool
}
func (value *APLValueCurrentRuneActive) GetBool(sim *Simulation) bool {
	return value.unit.RuneIsActive(value.runeSlot)
}

type APLValueCurrentRuneDeath struct {
	defaultAPLValueImpl
	unit     *Unit
	runeSlot int8
}

func (rot *APLRotation) newValueCurrentRuneDeath(config *proto.APLValueCurrentRuneDeath) APLValue {
	unit := rot.unit
	if !unit.HasRunicPowerBar() {
		return nil
	}
	return &APLValueCurrentRuneDeath{
		unit:     unit,
		runeSlot: int8(config.RuneSlot),
	}
}
func (value *APLValueCurrentRuneDeath) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeBool
}
func (value *APLValueCurrentRuneDeath) GetBool(sim *Simulation) bool {
	return value.unit.RuneIsDeath(int8(value.runeSlot))
}

type APLValueRuneCooldown struct {
	defaultAPLValueImpl
	unit     *Unit
	runeType proto.APLValueRuneType
}

func (rot *APLRotation) newValueRuneCooldown(config *proto.APLValueRuneCooldown) APLValue {
	unit := rot.unit
	if !unit.HasRunicPowerBar() {
		return nil
	}
	return &APLValueRuneCooldown{
		unit:     unit,
		runeType: config.RuneType,
	}
}
func (value *APLValueRuneCooldown) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeDuration
}
func (value *APLValueRuneCooldown) GetDuration(sim *Simulation) time.Duration {
	switch value.runeType {
	case proto.APLValueRuneType_RuneBlood:
		return value.unit.BloodRuneReadyAt(sim) - sim.CurrentTime
	case proto.APLValueRuneType_RuneFrost:
		return value.unit.FrostRuneReadyAt(sim) - sim.CurrentTime
	case proto.APLValueRuneType_RuneUnholy:
		return value.unit.UnholyRuneReadyAt(sim) - sim.CurrentTime
	}
	return 0
}

type APLValueNextRuneCooldown struct {
	defaultAPLValueImpl
	unit     *Unit
	runeType proto.APLValueRuneType
}

func (rot *APLRotation) newValueNextRuneCooldown(config *proto.APLValueNextRuneCooldown) APLValue {
	unit := rot.unit
	if !unit.HasRunicPowerBar() {
		return nil
	}
	return &APLValueNextRuneCooldown{
		unit:     unit,
		runeType: config.RuneType,
	}
}
func (value *APLValueNextRuneCooldown) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeDuration
}
func (value *APLValueNextRuneCooldown) GetDuration(sim *Simulation) time.Duration {
	switch value.runeType {
	case proto.APLValueRuneType_RuneBlood:
		return value.unit.SpentBloodRuneReadyAt() - sim.CurrentTime
	case proto.APLValueRuneType_RuneFrost:
		return value.unit.SpentFrostRuneReadyAt() - sim.CurrentTime
	case proto.APLValueRuneType_RuneUnholy:
		return value.unit.SpentUnholyRuneReadyAt() - sim.CurrentTime
	}
	return 0
}