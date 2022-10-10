package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/player"
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

func New(spell string) Spell {
    switch spell {
    case "Acid Splash":
        return AcidSplash{level: 0, casttime: "action", range: 60, components: "V,S", duration: 0, concentration: false}
    case "Alarm":
        return Alarm{level: 1, casttime: "minute", range: 30, components: "V,S,M", duration: 480, concentration: false}
    case "Animal Friendship":
        return AnimalFriendship{level: 1, casttime: "action", range: 30, components: "V,S,M", duration: 1440, concentration: false}
    case "Armor of Agathys":
        return ArmorOfAgathys{level: 1, casttime: "action", range: 0, components: "V,S,M", duration: 60, concentration: false}
    case "Arms of Hadar":
        return ArmsOfHadar{level: 1, casttime: "action", range: 0, components: "V,S", duration: 0, concentration: false}
    case "Bane":
        return Bane{level: 1, casttime: "action", range: 30, components: "V,S,M", duration: 1, concentration: true}
    case "Blade Ward":
        return BladeWard{level: 0, casttime: "action", range: 0, components: "V,S", duration: -1, concentration: false}
    case "Bless":
        return Bless{level: 1, casttime: "action", range: 30, components: "V,S,M", duration: 1, concentration: true}
    case "Burning Hands":
        return BurningHands{level: 1, casttime: "action", range: 0, components: "V,S", duration: 0, concentration: false}
    case "Charm Person":
        return CharmPerson{level: 1, casttime: "action", range: 30, components: "V,S", duration: 60, concentration: false}
    case "Chill Touch":
        return ChillTouch{level: 0, casttime: "action", range: 120, components: "V,S", duration: -1}
    case "Chromatic Orb":
        return ChromaticOrb{level: 1, casttime: "action", range: 90, components: "V,S,M", duration: 0, concentration: false}
    case "Color Spray":
        return ColorSpray{level: 1, casttime: "action", range: 0, components: "V,S,M", duration: -1, concentration: false}
    case "Command":
        return Command{level: 1, casttime: "action", range: 60, components: "V", duration: -1, concentration: false}
    case "Compelled Duel":
        return CompelledDuel{level: 1, casttime: "bonus", range: 30, components: "V", duration: 1, concentration: true}
    case "Comprehend Languages":
        return ComprehendLanguages{level: 1, casttime: "action", range: 0, components: "V,S,M", duration: 60, concentration: false}
    case "Create or Destroy Water":
        return CreateOrDestroyWater{level: 1, casttime: "action", range: 30, components: "V,S,M", duration: 0, concentration: false}
    case "Cure Wounds":
        return CureWounds{level: 1, casttime: "action", range: 5, components: "V,S", duration: 0, concentration: false}
    case "Dancing Lights":
        return DancingLights{level: 0, casttime: "action", range: 120, components: "V,S,M", duration: 1, concentration: true}
    case "Detect Evil and Good":
        return DetectEvilAndGood{level: 1, casttime: "action", range: 0, components: "V,S", duration: 10, concentration: true}
    case "Detect Magic":
        return DetectMagic{level: 1, casttime: "action", range: 0, components: "V,S", duration: 10, concentration: true}
    case "Detect Poison and Disease":
        return DetectPoisonAndDisease{level: 1, casttime: "action", range: 0, components: "V,S,M", duration: 10, concentration: true}
    case "Disguise Self":
        return DisguiseSelf{level: 1, casttime: "action", range: 0, components: "V,S", duration: 60, concentration: false}
    case "Dissonant Whispers":
        return DissonantWhispers{level: 1, casttime: "action", range: 60, components: "V", duration: 0, concentration: false}
    case "Divine Favor":
        return DivineFavor{level: 1, casttime: "bonus", range: 0, components: "V,S", duration: 1, concentration: true}
    case "Druidcraft":
        return Druidcraft{level: 0, casttime: "action", range: 30, components: "V,S", duration: 0, concentration: false}
    case "Eldritch Blast":
        return EldritchBlast{level: 0, casttime: "action", range: 120, components: "V,S", duration: 0, concentration: false}
    case "Ensnaring Strike":
        return EnsnaringStrike{level: 1, casttime: "bonus", range: 0, components: "V", duration: 1, concentration: true}
    case "Entangle":
        return Entangle{level: 1, casttime: "action", range: 90, components: "V,S", duration: 1, concentration: true}
    case "Expeditious Retreat":
        return ExpeditiousRetreat{level: 1, casttime: "bonus", range: 0, components: "V,S", duration: 10, concentration: true}
    case "Faerie Fire":
        return FaerieFire{level: 1, casttime: "action", range: 60, components: "V", duration: 1, concentration: true}
    case "False Life":
        return FalseLife{level: 1, casttime: "action", range: 0, components: "V,S,M", duration: 60, concentration: false}
    case "Feather Fall":
        return FeatherFall{level: 1, casttime: "reaction", range: 60, components: "V,M", duration: 1, concentration: false}
    case "Find Familiar":
        return FindFamiliar{level: 1, casttime: "hour", range: 10, components: "V,S,M", duration: 0, concentration: false}
    case "Fire Bolt":
        return FireBolt{level: 0, casttime: "action", range: 120, components: "V,S", duration: 0, concentration: false}
    case "Fog Cloud":
        return FogCloud{level: 1, casttime: "action", range: 120, components: "V,S", duration: 60, concentration: true}
    case "Friends":
        return Friends{level: 0, casttime: "action", range: 0, components: "S,M", duration: 1, concentration: true}
    case "Goodberry":
        return Goodberry{level: 1, casttime: "action", range: 5, components: "V,S,M", duration: 0, concentration: false}
    case "Grease":
        return Grease{level: 1, casttime: "action", range: 60, components: "V,S,M", duration: 1, concentration: false}
    case "Guidance":
        return Guidance{level: 0, casttime: "action", range: 5, components: "V,S", duration: 1, concentration: true}
    case "Guiding Bolt":
        return GuidingBolt{level: 1, casttime: "action", range: 120, components: "V,S", duration: -1, concentration: false}
    case "Hail of Thorns":
        return HailOfThorns{level: 1, casttime: "bonus", range: 0, components: "V", duration: 1, concentration: true}
    case "Healing Word":
        return HealingWord{level: 1, casttime: "bonus", range: 60, components: "V", duration: 0, concentration: false}
    case "Hellish Rebuke":
        return HellishRebuke{level: 1, casttime: "reaction", range: 60, components: "V,S", duration: 0, concentration: false}
    case "Heroism":
        return Heroism{level: 1, casttime: "action", range: 5, components: "V,S", duration: 1, concentration: true}
    case "Hex":
        return Hex{level: 1, casttime: "bonus", range: 90, components: "V,S,M", duration: 60, concentration: true}
    case "Hunter's Mark":
        return HuntersMark{level: 1, casttime: "bonus", range: 90, components: "V", duration: 60, concentration: true}
    case "Identify":
        return Identify{level: 1, casttime: "minute", range: 5, components: "V,S,M", duration: 0, concentration: false}
    case "Illusory Script":
        return IllusoryScript{level: 1, casttime: "minute", range: 5, components: "S,M", duration: 14400, concentration: false}
    case "Inflict Wounds":
        return InflictWounds{level: 1, casttime: "action", range: 5, components: "V,S", duration: 0, concentration: false}
    case "Jump":
        return Jump{level: 1, casttime: "action", range: 5, components: "V,S,M", duration: 1, concentration: false}
    case "Light":
        return Light{level: 0, casttime: "action", range: 5, components: "V,M", duration: 60, concentration: false}
    case "Longstrider":
        return Longstrider{level: 1, casttime: "action", range: 5, components: "V,S,M", duration: 60, concentration: false}
    case "Mage Armor":
        return MageArmor{level: 1, casttime: "action", range: 5, components: "V,S,M", duration: 480, concentration: false}
    case "Mage Hand":
        return MageHand{level: 0, casttime: "action", range: 30, components: "V,S", duration: 1, concentration: false}
    case "Magic Missile":
        return MagicMissile{level: 1, casttime: "action", range: 120, components: "V,S", duration: 0, concentration: false}
    case "Mending":
        return Mending{level: 0, casttime: "minute", range: 5, components: "V,S,M", duration: 0, concentration: false}
    case "Message":
        return Message{level: 0, casttime: "action", range: 120, components: "V,S,M", duration: -1, concentration: false}
    case "Minor Illusion":
        return MinorIllusion{level: 0, casttime: "action", range: 30, components: "S,M", duration: 1, concentration: false}
    case "Poison Spray":
        return PoisonSpray{level: 0, casttime: "action", range: 10, components: "V,S", duration: 0, concentration: false}
    case "Prestidigitation":
        return Prestidigitation{level: 0, casttime: "action", range: 10, components: "V,S", duration: 60, concentration: false}
    case "Produce Flame":
        return ProduceFlame{level: 0, casttime: "action", range: 0, components: "V,S", duration: 10, concentration: false}
    case "Protection from Good and Evil":
        return ProtectionFromGoodAndEvil{level: 1, casttime: "action", range: 5, components: "V,S,M", duration: 10, concentration: true}
    case "Purify Food and Drink":
        return PurifyFoodAndDrink{level: 1, casttime: "action", range: 10, components: "V,S", duration: 0, concentration: false}
    case "Ray of Frost":
        return RayOfFrost{level: 0, casttime: "action", range: 60, components: "V,S", duration: 0, concentration: false}
    case "Ray of Sickness":
        return RayOfSickness{level: 1, casttime: "action", range: 60, components: "V,S", duration: 0, concentration: false}
    case "Resistance":
        return Resistance{level: 0, casttime: "action", range: 5, components: "V,S,M", duration: 1, concentration: true}
    case "Sacred Flame":
        return SacredFlame{level: 0, casttime: "action", range: 60, components: "V,S", duration: 0, concentration: false}
    case "Sanctuary":
        return Sanctuary{level: 1, casttime: "bonus", range: 30, components: "V,S,M", duration: 1, concentration: false}
    case "Searing Smite":
        return SearingSmite{level: 1, casttime: "bonus", range: 0, components: "V", duration: 1, concentration: true}
    case "Shield":
        return Shield{level: 1, casttime: "reaction", range: 0, components: "V,S", duration: -1, concentration: false}
    case "Shield of Faith":
        return ShieldOfFaith{level: 1, casttime: "bonus", range: 60, components: "V,S,M", duration: 10, concentration: true}
    case "Shillelagh":
        return Shillelagh{level: 0, casttime: "bonus", range: 5, components: "V,S,M", duration: 1, concentration: false}
    case "Shocking Grasp":
        return ShockingGrasp{level: 0, casttime: "action", range: 5, components: "V,S", duration: 0, concentration: false}
    case "Silent Image":
        return SilentImage{level: 1, casttime: "action", range: 60, components: "V,S,M", duration: 10, concentration: true}
    case "Sleep":
        return Sleep{level: 1, casttime: "action", range: 90, components: "V,S,M", duration: 1, concentration: false}
    case "Spare the Dying":
        return SpareTheDying{level: 0, casttime: "action", range: 5, components: "V,S", duration: 0, concentration: false}
    case "Speak with Animals":
        return SpeakWithAnimals{level: 1, casttime: "action", range: 0, components: "V,S", duration: 10, concentration: false}
    case "Tasha's Hideous Laughter":
        return TashasHideousLaughter{level: 1, casttime: "action", range: 30, components: "V,S,M", duration: 1, concentration: true}
    case "Tenser's Floating Disk":
        return TensersFloatingDisk{level: 1, casttime: "action", range: 30, components: "V,S,M", duration: 60, concentration: false}
    case "Thaumaturgy":
        return Thaumaturgy{level: 0, casttime: "action", range: 30, components: "V", duration: 1, concentration: false}
    case "Thorn Whip":
        return ThornWhip{level: 0, casttime: "action", range: 30, components: "V,S,M", duration: 0, concentration: false}
    case "Thunderous Smite":
        return ThunderousSmite{level: 1, casttime: "bonus", range: 0, components: "V", duration: 1, concentration: true}
    case "Thunderwave":
        return Thunderwave{level: 1, casttime: "action", range: 0, components: "V,S", duration: 0, concentration: false}
    case "True Strike":
        return TrueStrike{level: 0, casttime: "action", range: 30, components: "S", duration: -1, concentration: true}
    case "Unseen Servant":
        return UnseenServant{level: 1, casttime: "action", range: 60, components: "V,S,M", duration: 60, concentration: false}
    case "Vicious Mockery":
        return ViciousMockery{level: 0, casttime: "action", range: 60, components: "V", duration: 0, concentration: false}
    case "Witch Bolt":
        return WitchBolt{level: 1, casttime: "action", range: 30, components: "V,S,M", duration: 1, concentration: true}
    case "Wrathful Smite":
        return WrathfulSmite{level: 1, casttime: "bonus", range: 0, components: "V", duration: 1, concentration: true}
    default:
        return errors.New(fmt.Sprintf("%s is not a valid spell", spell))
    }
}
