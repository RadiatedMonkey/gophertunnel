package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/RadiatedMonkey/gophertunnel/minecraft/protocol"
)

// noinspection SpellCheckingInspection
const (
	SoundEventItemUseOn = iota
	SoundEventHit
	SoundEventStep
	SoundEventFly
	SoundEventJump
	SoundEventBreak
	SoundEventPlace
	SoundEventHeavyStep
	SoundEventGallop
	SoundEventFall
	SoundEventAmbient
	SoundEventAmbientBaby
	SoundEventAmbientInWater
	SoundEventBreathe
	SoundEventDeath
	SoundEventDeathInWater
	SoundEventDeathToZombie
	SoundEventHurt
	SoundEventHurtInWater
	SoundEventMad
	SoundEventBoost
	SoundEventBow
	SoundEventSquishBig
	SoundEventSquishSmall
	SoundEventFallBig
	SoundEventFallSmall
	SoundEventSplash
	SoundEventFizz
	SoundEventFlap
	SoundEventSwim
	SoundEventDrink
	SoundEventEat
	SoundEventTakeoff
	SoundEventShake
	SoundEventPlop
	SoundEventLand
	SoundEventSaddle
	SoundEventArmor
	SoundEventArmorPlace
	SoundEventAddChest
	SoundEventThrow
	SoundEventAttack
	SoundEventAttackNoDamage
	SoundEventAttackStrong
	SoundEventWarn
	SoundEventShear
	SoundEventMilk
	SoundEventThunder
	SoundEventExplode
	SoundEventFire
	SoundEventIgnite
	SoundEventFuse
	SoundEventStare
	SoundEventSpawn
	SoundEventShoot
	SoundEventBreakBlock
	SoundEventLaunch
	SoundEventBlast
	SoundEventLargeBlast
	SoundEventTwinkle
	SoundEventRemedy
	SoundEventUnfect
	SoundEventLevelUp
	SoundEventBowHit
	SoundEventBulletHit
	SoundEventExtinguishFire
	SoundEventItemFizz
	SoundEventChestOpen
	SoundEventChestClosed
	SoundEventShulkerBoxOpen
	SoundEventShulkerBoxClosed
	SoundEventEnderChestOpen
	SoundEventEnderChestClosed
	SoundEventPowerOn
	SoundEventPowerOff
	SoundEventAttach
	SoundEventDetach
	SoundEventDeny
	SoundEventTripod
	SoundEventPop
	SoundEventDropSlot
	SoundEventNote
	SoundEventThorns
	SoundEventPistonIn
	SoundEventPistonOut
	SoundEventPortal
	SoundEventWater
	SoundEventLavaPop
	SoundEventLava
	SoundEventBurp
	SoundEventBucketFillWater
	SoundEventBucketFillLava
	SoundEventBucketEmptyWater
	SoundEventBucketEmptyLava
	SoundEventEquipChain
	SoundEventEquipDiamond
	SoundEventEquipGeneric
	SoundEventEquipGold
	SoundEventEquipIron
	SoundEventEquipLeather
	SoundEventEquipElytra
	SoundEventRecord13
	SoundEventRecordCat
	SoundEventRecordBlocks
	SoundEventRecordChirp
	SoundEventRecordFar
	SoundEventRecordMall
	SoundEventRecordMellohi
	SoundEventRecordStal
	SoundEventRecordStrad
	SoundEventRecordWard
	SoundEventRecord11
	SoundEventRecordWait
	SoundEventRecordNull
	SoundEventFlop
	SoundEventGuardianCurse
	SoundEventMobWarning
	SoundEventMobWarningBaby
	SoundEventTeleport
	SoundEventShulkerOpen
	SoundEventShulkerClose
	SoundEventHaggle
	SoundEventHaggleYes
	SoundEventHaggleNo
	SoundEventHaggleIdle
	SoundEventChorusGrow
	SoundEventChorusDeath
	SoundEventGlass
	SoundEventPotionBrewed
	SoundEventCastSpell
	SoundEventPrepareAttackSpell
	SoundEventPrepareSummon
	SoundEventPrepareWololo
	SoundEventFang
	SoundEventCharge
	SoundEventTakePicture
	SoundEventPlaceLeashKnot
	SoundEventBreakLeashKnot
	SoundEventAmbientGrowl
	SoundEventAmbientWhine
	SoundEventAmbientPant
	SoundEventAmbientPurr
	SoundEventAmbientPurreow
	SoundEventDeathMinVolume
	SoundEventDeathMidVolume
	SoundEventImitateBlaze
	SoundEventImitateCaveSpider
	SoundEventImitateCreeper
	SoundEventImitateElderGuardian
	SoundEventImitateEnderDragon
	SoundEventImitateEnderman
	SoundEventImitateEndermite
	SoundEventImitateEvocationIllager
	SoundEventImitateGhast
	SoundEventImitateHusk
	SoundEventImitateIllusionIllager
	SoundEventImitateMagmaCube
	SoundEventImitatePolarBear
	SoundEventImitateShulker
	SoundEventImitateSilverfish
	SoundEventImitateSkeleton
	SoundEventImitateSlime
	SoundEventImitateSpider
	SoundEventImitateStray
	SoundEventImitateVex
	SoundEventImitateVindicationIllager
	SoundEventImitateWitch
	SoundEventImitateWither
	SoundEventImitateWitherSkeleton
	SoundEventImitateWolf
	SoundEventImitateZombie
	SoundEventImitateZombiePigman
	SoundEventImitateZombieVillager
	SoundEventEnderEyePlaced
	SoundEventEndPortalCreated
	SoundEventAnvilUse
	SoundEventBottleDragonBreath
	SoundEventPortalTravel
	SoundEventTridentHit
	SoundEventTridentReturn
	SoundEventTridentRiptide1
	SoundEventTridentRiptide2
	SoundEventTridentRiptide3
	SoundEventTridentThrow
	SoundEventTridentThunder
	SoundEventTridentHitGround
	SoundEventDefault
	SoundEventFletchingTableUse
	SoundEventElemConstructOpen
	SoundEventIceBombHit
	SoundEventBalloonPop
	SoundEventLtReactionIceBomb
	SoundEventLtReactionBleach
	SoundEventLtReactionElephantToothpaste
	SoundEventLtReactionElephantToothpaste2
	SoundEventLtReactionGlowStick
	SoundEventLtReactionGlowStick2
	SoundEventLtReactionLuminol
	SoundEventLtReactionSalt
	SoundEventLtReactionFertilizer
	SoundEventLtReactionFireball
	SoundEventLtReactionMagnesiumSalt
	SoundEventLtReactionMiscFire
	SoundEventLtReactionFire
	SoundEventLtReactionMiscExplosion
	SoundEventLtReactionMiscMystical
	SoundEventLtReactionMiscMystical2
	SoundEventLtReactionProduct
	SoundEventSparklerUse
	SoundEventGlowStickUse
	SoundEventSparklerActive
	SoundEventConvertToDrowned
	SoundEventBucketFillFish
	SoundEventBucketEmptyFish
	SoundEventBubbleColumnUpwards
	SoundEventBubbleColumnDownwards
	SoundEventBubblePop
	SoundEventBubbleUpInside
	SoundEventBubbleDownInside
	SoundEventHurtBaby
	SoundEventDeathBaby
	SoundEventStepBaby
	SoundEventSpawnBaby
	SoundEventBorn
	SoundEventTurtleEggBreak
	SoundEventTurtleEggCrack
	SoundEventTurtleEggHatched
	SoundEventLayEgg
	SoundEventTurtleEggAttacked
	SoundEventBeaconActivate
	SoundEventBeaconAmbient
	SoundEventBeaconDeactivate
	SoundEventBeaconPower
	SoundEventConduitActivate
	SoundEventConduitAmbient
	SoundEventConduitAttack
	SoundEventConduitDeactivate
	SoundEventConduitShort
	SoundEventSwoop
	SoundEventBambooSaplingPlace
	SoundEventPreSneeze
	SoundEventSneeze
	SoundEventAmbientTame
	SoundEventScared
	SoundEventScaffoldingClimb
	SoundEventCrossbowLoadingStart
	SoundEventCrossbowLoadingMiddle
	SoundEventCrossbowLoadingEnd
	SoundEventCrossbowShoot
	SoundEventCrossbowQuickChargeStart
	SoundEventCrossbowQuickChargeMiddle
	SoundEventCrossbowQuickChargeEnd
	SoundEventAmbientAggressive
	SoundEventAmbientWorried
	SoundEventCantBreed
	SoundEventShieldBlock
	SoundEventLecternBookPlace
	SoundEventGrindstoneUse
	SoundEventBell
	SoundEventCampfireCrackle
	SoundEventRoar
	SoundEventStun
	SoundEventSweetBerryBushHurt
	SoundEventSweetBerryBushPick
	SoundEventCartographyTableUse
	SoundEventStonecutterUse
	SoundEventComposterEmpty
	SoundEventComposterFill
	SoundEventComposterFillLayer
	SoundEventComposterReady
	SoundEventBarrelOpen
	SoundEventBarrelClose
	SoundEventRaidHorn
	SoundEventLoomUse
	SoundEventAmbientInRaid
	SoundEventUicartographyTableUse
	SoundEventUistonecutterUse
	SoundEventUiloomUse
	SoundEventSmokerUse
	SoundEventBlastFurnaceUse
	SoundEventSmithingTableUse
	SoundEventScreech
	SoundEventSleep
	SoundEventFurnaceUse
	SoundEventMooshroomConvert
	SoundEventMilkSuspiciously
	SoundEventCelebrate
	SoundEventJumpPrevent
	SoundEventAmbientPollinate
	SoundEventBeehiveDrip
	SoundEventBeehiveEnter
	SoundEventBeehiveExit
	SoundEventBeehiveWork
	SoundEventBeehiveShear
	SoundEventHoneybottleDrink
	SoundEventAmbientCave
	SoundEventRetreat
	SoundEventConvertToZombified
	SoundEventAdmire
	SoundEventStepLava
	SoundEventTempt
	SoundEventPanic
	SoundEventAngry
	SoundEventAmbientMoodWarpedForest
	SoundEventAmbientMoodSoulsandValley
	SoundEventAmbientMoodNetherWastes
	SoundEventAmbientMoodBasaltDeltas
	SoundEventAmbientMoodCrimsonForest
	SoundEventRespawnAnchorCharge
	SoundEventRespawnAnchorDeplete
	SoundEventRespawnAnchorSetSpawn
	SoundEventRespawnAnchorAmbient
	SoundEventSoulEscapeQuiet
	SoundEventSoulEscapeLoud
	SoundEventRecordPigstep
	SoundEventLinkCompassToLodestone
	SoundEventUseSmithingTable
	SoundEventEquipNetherite
	SoundEventAmbientLoopWarpedForest
	SoundEventAmbientLoopSoulsandValley
	SoundEventAmbientLoopNetherWastes
	SoundEventAmbientLoopBasaltDeltas
	SoundEventAmbientLoopCrimsonForest
	SoundEventAmbientAdditionWarpedForest
	SoundEventAmbientAdditionSoulsandValley
	SoundEventAmbientAdditionNetherWastes
	SoundEventAmbientAdditionBasaltDeltas
	SoundEventAmbientAdditionCrimsonForest
	SoundEventSculkSensorPowerOn
	SoundEventSculkSensorPowerOff
	SoundEventBucketFillPowderSnow
	SoundEventBucketEmptyPowderSnow
	SoundEventPointedDripstoneCauldronDripWater
	SoundEventPointedDripstoneCauldronDripLava
	SoundEventPointedDripstoneDripWater
	SoundEventPointedDripstoneDripLava
	SoundEventCaveVinesPickBerries
	SoundEventBigDripleafTiltDown
	SoundEventBigDripleafTiltUp
	SoundEventCopperWaxOn
	SoundEventCopperWaxOff
	SoundEventScrape
	SoundEventPlayerHurtDrown
	SoundEventPlayerHurtOnFire
	SoundEventPlayerHurtFreeze
	SoundEventUseSpyglass
	SoundEventStopUsingSpyglass
	SoundEventAmethystBlockChime
	SoundEventAmbientScreamer
	SoundEventHurtScreamer
	SoundEventDeathScreamer
	SoundEventMilkScreamer
	SoundEventJumpToBlock
	SoundEventPreRam
	SoundEventPreRamScreamer
	SoundEventRamImpact
	SoundEventRamImpactScreamer
	SoundEventSquidInkSquirt
	SoundEventGlowSquidInkSquirt
	SoundEventConvertToStray
	SoundEventCakeAddCandle
	SoundEventExtinguishCandle
	SoundEventAmbientCandle
	SoundEventBlockClick
	SoundEventBlockClickFail
	SoundEventSculkCatalystBloom
	SoundEventSculkShriekerShriek
	SoundEventWardenNearbyClose
	SoundEventWardenNearbyCloser
	SoundEventWardenNearbyClosest
	SoundEventWardenSlightlyAngry
	SoundEventRecordOtherside
	SoundEventTongue
	SoundEventCrackIronGolem
	SoundEventRepairIronGolem
	SoundEventListening
	SoundEventHeartbeat
	SoundEventHornBreak
	_
	SoundEventSculkSpread
	SoundEventSculkCharge
	SoundEventSculkSensorPlace
	SoundEventSculkShriekerPlace
	SoundEventGoatCall0
	SoundEventGoatCall1
	SoundEventGoatCall2
	SoundEventGoatCall3
	SoundEventGoatCall4
	SoundEventGoatCall5
	SoundEventGoatCall6
	SoundEventGoatCall7
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	SoundEventImitateWarden
	SoundEventListeningAngry
	SoundEventItemGiven
	SoundEventItemTaken
	SoundEventDisappeared
	SoundEventReappeared
	SoundEventDrinkMilk
	SoundEventFrogspawnHatched
	SoundEventLaySpawn
	SoundEventFrogspawnBreak
	SoundEventSonicBoom
	SoundEventSonicCharge
	SoundeventItemThrown
	SoundEventRecord5
	SoundEventConvertToFrog
	SoundEventRecordPlaying
	SoundEventEnchantingTableUse
	SoundEventStepSand
	SoundEventDashReady
	SoundEventBundleDropContents
	SoundEventBundleInsert
	SoundEventBundleRemoveOne
	SoundEventPressurePlateClickOff
	SoundEventPressurePlateClickOn
	SoundEventButtonClickOff
	SoundEventButtonClickOn
	SoundEventDoorOpen
	SoundEventDoorClose
	SoundEventTrapdoorOpen
	SoundEventTrapdoorClose
	SoundEventFenceGateOpen
	SoundEventFenceGateClose
	SoundEventInsert
	SoundEventPickup
	SoundEventInsertEnchanted
	SoundEventPickupEnchanted
)

// LevelSoundEvent is sent by the server to make any kind of built-in sound heard to a player. It is sent to,
// for example, play a stepping sound or a shear sound. The packet is also sent by the client, in which case
// it could be forwarded by the server to the other players online. If possible, the packets from the client
// should be ignored however, and the server should play them on its own accord.
type LevelSoundEvent struct {
	// SoundType is the type of the sound to play. It is one of the constants above. Some of the sound types
	// require additional data, which is set in the EventData field.
	SoundType uint32
	// Position is the position of the sound event. The player will be able to hear the direction of the sound
	// based on what position is sent here.
	Position mgl32.Vec3
	// ExtraData is a packed integer that some sound types use to provide extra data. An example of this is
	// the note sound, which is composed of a pitch and an instrument type.
	ExtraData int32
	// EntityType is the string entity type of the entity that emitted the sound, for example
	// 'minecraft:skeleton'. Some sound types use this entity type for additional data.
	EntityType string
	// BabyMob specifies if the sound should be that of a baby mob. It is most notably used for parrot
	// imitations, which will change based on if this field is set to true or not.
	BabyMob bool
	// DisableRelativeVolume specifies if the sound should be played relatively or not. If set to true, the
	// sound will have full volume, regardless of where the Position is, whereas if set to false, the sound's
	// volume will be based on the distance to Position.
	DisableRelativeVolume bool
}

// ID ...
func (*LevelSoundEvent) ID() uint32 {
	return IDLevelSoundEvent
}

// Marshal ...
func (pk *LevelSoundEvent) Marshal(w *protocol.Writer) {
	w.Varuint32(&pk.SoundType)
	w.Vec3(&pk.Position)
	w.Varint32(&pk.ExtraData)
	w.String(&pk.EntityType)
	w.Bool(&pk.BabyMob)
	w.Bool(&pk.DisableRelativeVolume)
}

// Unmarshal ...
func (pk *LevelSoundEvent) Unmarshal(r *protocol.Reader) {
	r.Varuint32(&pk.SoundType)
	r.Vec3(&pk.Position)
	r.Varint32(&pk.ExtraData)
	r.String(&pk.EntityType)
	r.Bool(&pk.BabyMob)
	r.Bool(&pk.DisableRelativeVolume)
}
