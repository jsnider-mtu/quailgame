package spells

import (
    "errors"
    "fmt"

    "github.com/hajimehoshi/ebiten/v2"
)

type Spell interface {
    PrettyPrint() string
    GetLevel() int
    GetCastTime() string
    GetRange() int
    GetComponents() string
    GetDuration() int
    GetConcentration() bool
    Animate(*ebiten.Image, [2]int)
}

func New(spell string) (Spell, error) {
    switch spell {
    case "Acid Splash":
        return AcidSplash{level: 0, casttime: "action", spellrange: 60, components: "V,S", duration: 0, concentration: false}, nil
    case "Armor of Agathys":
        return ArmorOfAgathys{level: 1, casttime: "action", spellrange: 0, components: "V,S,M", duration: 60, concentration: false}, nil
    case "Arms of Hadar":
        return ArmsOfHadar{level: 1, casttime: "action", spellrange: 0, components: "V,S", duration: 0, concentration: false}, nil
    case "Bane":
        return Bane{level: 1, casttime: "action", spellrange: 30, components: "V,S,M", duration: 1, concentration: true}, nil
    case "Blade Ward":
        return BladeWard{level: 0, casttime: "action", spellrange: 0, components: "V,S", duration: -1, concentration: false}, nil
    case "Bless":
        return Bless{level: 1, casttime: "action", spellrange: 30, components: "V,S,M", duration: 1, concentration: true}, nil
    case "Burning Hands":
        return BurningHands{level: 1, casttime: "action", spellrange: 0, components: "V,S", duration: 0, concentration: false}, nil
    case "Chill Touch":
        return ChillTouch{level: 0, casttime: "action", spellrange: 120, components: "V,S", duration: -1}, nil
    case "Chromatic Orb":
        return ChromaticOrb{level: 1, casttime: "action", spellrange: 90, components: "V,S,M", duration: 0, concentration: false}, nil
    case "Color Spray":
        return ColorSpray{level: 1, casttime: "action", spellrange: 0, components: "V,S,M", duration: -1, concentration: false}, nil
    case "Comprehend Languages":
        return ComprehendLanguages{level: 1, casttime: "action", spellrange: 0, components: "V,S,M", duration: 60, concentration: false}, nil
    case "Cure Wounds":
        return CureWounds{level: 1, casttime: "action", spellrange: 5, components: "V,S", duration: 0, concentration: false}, nil
    case "Dissonant Whispers":
        return DissonantWhispers{level: 1, casttime: "action", spellrange: 60, components: "V", duration: 0, concentration: false}, nil
    case "Divine Favor":
        return DivineFavor{level: 1, casttime: "bonus", spellrange: 0, components: "V,S", duration: 1, concentration: true}, nil
    case "Eldritch Blast":
        return EldritchBlast{level: 0, casttime: "action", spellrange: 120, components: "V,S", duration: 0, concentration: false}, nil
    case "Ensnaring Strike":
        return EnsnaringStrike{level: 1, casttime: "bonus", spellrange: 0, components: "V", duration: 1, concentration: true}, nil
    case "Entangle":
        return Entangle{level: 1, casttime: "action", spellrange: 90, components: "V,S", duration: 1, concentration: true}, nil
    case "Expeditious Retreat":
        return ExpeditiousRetreat{level: 1, casttime: "bonus", spellrange: 0, components: "V,S", duration: 10, concentration: true}, nil
    case "Faerie Fire":
        return FaerieFire{level: 1, casttime: "action", spellrange: 60, components: "V", duration: 1, concentration: true}, nil
    case "False Life":
        return FalseLife{level: 1, casttime: "action", spellrange: 0, components: "V,S,M", duration: 60, concentration: false}, nil
    case "Fire Bolt":
        return FireBolt{level: 0, casttime: "action", spellrange: 120, components: "V,S", duration: 0, concentration: false}, nil
    case "Fog Cloud":
        return FogCloud{level: 1, casttime: "action", spellrange: 120, components: "V,S", duration: 60, concentration: true}, nil
    case "Grease":
        return Grease{level: 1, casttime: "action", spellrange: 60, components: "V,S,M", duration: 1, concentration: false}, nil
    case "Guiding Bolt":
        return GuidingBolt{level: 1, casttime: "action", spellrange: 120, components: "V,S", duration: -1, concentration: false}, nil
    case "Hail of Thorns":
        return HailOfThorns{level: 1, casttime: "bonus", spellrange: 0, components: "V", duration: 1, concentration: true}, nil
    case "Healing Word":
        return HealingWord{level: 1, casttime: "bonus", spellrange: 60, components: "V", duration: 0, concentration: false}, nil
    case "Hellish Rebuke":
        return HellishRebuke{level: 1, casttime: "reaction", spellrange: 60, components: "V,S", duration: 0, concentration: false}, nil
    case "Heroism":
        return Heroism{level: 1, casttime: "action", spellrange: 5, components: "V,S", duration: 1, concentration: true}, nil
    case "Hex":
        return Hex{level: 1, casttime: "bonus", spellrange: 90, components: "V,S,M", duration: 60, concentration: true}, nil
    case "Hunter's Mark":
        return HuntersMark{level: 1, casttime: "bonus", spellrange: 90, components: "V", duration: 60, concentration: true}, nil
    case "Inflict Wounds":
        return InflictWounds{level: 1, casttime: "action", spellrange: 5, components: "V,S", duration: 0, concentration: false}, nil
    case "Longstrider":
        return Longstrider{level: 1, casttime: "action", spellrange: 5, components: "V,S,M", duration: 60, concentration: false}, nil
    case "Mage Armor":
        return MageArmor{level: 1, casttime: "action", spellrange: 5, components: "V,S,M", duration: 480, concentration: false}, nil
    case "Magic Missile":
        return MagicMissile{level: 1, casttime: "action", spellrange: 120, components: "V,S", duration: 0, concentration: false}, nil
    case "Poison Spray":
        return PoisonSpray{level: 0, casttime: "action", spellrange: 10, components: "V,S", duration: 0, concentration: false}, nil
    case "Produce Flame":
        return ProduceFlame{level: 0, casttime: "action", spellrange: 0, components: "V,S", duration: 10, concentration: false}, nil
    case "Protection from Good and Evil":
        return ProtectionFromGoodAndEvil{level: 1, casttime: "action", spellrange: 5, components: "V,S,M", duration: 10, concentration: true}, nil
    case "Ray of Frost":
        return RayOfFrost{level: 0, casttime: "action", spellrange: 60, components: "V,S", duration: 0, concentration: false}, nil
    case "Ray of Sickness":
        return RayOfSickness{level: 1, casttime: "action", spellrange: 60, components: "V,S", duration: 0, concentration: false}, nil
    case "Resistance":
        return Resistance{level: 0, casttime: "action", spellrange: 5, components: "V,S,M", duration: 1, concentration: true}, nil
    case "Sacred Flame":
        return SacredFlame{level: 0, casttime: "action", spellrange: 60, components: "V,S", duration: 0, concentration: false}, nil
    case "Sanctuary":
        return Sanctuary{level: 1, casttime: "bonus", spellrange: 30, components: "V,S,M", duration: 1, concentration: false}, nil
    case "Searing Smite":
        return SearingSmite{level: 1, casttime: "bonus", spellrange: 0, components: "V", duration: 1, concentration: true}, nil
    case "Shield":
        return Shield{level: 1, casttime: "reaction", spellrange: 0, components: "V,S", duration: -1, concentration: false}, nil
    case "Shield of Faith":
        return ShieldOfFaith{level: 1, casttime: "bonus", spellrange: 60, components: "V,S,M", duration: 10, concentration: true}, nil
    case "Shillelagh":
        return Shillelagh{level: 0, casttime: "bonus", spellrange: 5, components: "V,S,M", duration: 1, concentration: false}, nil
    case "Shocking Grasp":
        return ShockingGrasp{level: 0, casttime: "action", spellrange: 5, components: "V,S", duration: 0, concentration: false}, nil
    case "Sleep":
        return Sleep{level: 1, casttime: "action", spellrange: 90, components: "V,S,M", duration: 1, concentration: false}, nil
    case "Spare the Dying":
        return SpareTheDying{level: 0, casttime: "action", spellrange: 5, components: "V,S", duration: 0, concentration: false}, nil
    case "Speak with Animals":
        return SpeakWithAnimals{level: 1, casttime: "action", spellrange: 0, components: "V,S", duration: 10, concentration: false}, nil
    case "Tasha's Hideous Laughter":
        return TashasHideousLaughter{level: 1, casttime: "action", spellrange: 30, components: "V,S,M", duration: 1, concentration: true}, nil
    case "Thorn Whip":
        return ThornWhip{level: 0, casttime: "action", spellrange: 30, components: "V,S,M", duration: 0, concentration: false}, nil
    case "Thunderous Smite":
        return ThunderousSmite{level: 1, casttime: "bonus", spellrange: 0, components: "V", duration: 1, concentration: true}, nil
    case "Thunderwave":
        return Thunderwave{level: 1, casttime: "action", spellrange: 0, components: "V,S", duration: 0, concentration: false}, nil
    case "True Strike":
        return TrueStrike{level: 0, casttime: "action", spellrange: 30, components: "S", duration: -1, concentration: true}, nil
    case "Vicious Mockery":
        return ViciousMockery{level: 0, casttime: "action", spellrange: 60, components: "V", duration: 0, concentration: false}, nil
    case "Witch Bolt":
        return WitchBolt{level: 1, casttime: "action", spellrange: 30, components: "V,S,M", duration: 1, concentration: true}, nil
    case "Wrathful Smite":
        return WrathfulSmite{level: 1, casttime: "bonus", spellrange: 0, components: "V", duration: 1, concentration: true}, nil
    default:
        return nil, errors.New(fmt.Sprintf("%s is not a valid spell", spell))
    }
}
