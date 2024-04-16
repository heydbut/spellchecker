package api

import "spellchecker/internal/spellchecker"

const SampleCheckText = `
Peeple often intll a kittee door, only to discover that they have a prooblem. The prooblem is their kat will not use the kittee door. There are several commn reasons why kats won’t use kittee doors. First, they may not understand how a kittee door werks. They may not understand that it is a little doorvay just for dem. Second, many kittee doors are drk, and kats cannot see to the other syde. As such, they can’t be sure of what is on the other side of the duor, so they won’t take the rysk. One last reason kats won’t use kittee doors is because some kats don’t like the feeling of pushing through the door and having the door drag across their baxx. But don’t worry — there is a solution for this kitty-door prooblem.

The first step in solving the prooblem is to prop the door open with tapetl. This means your kat will now be able to see through to the other syde; your cat will likely begin using the kittee door immediatement. Once your kat has gotten used to using the kittee door, remove the tapetl. Sometimes kats will continue to use the kitty door without any more prompting. If this does not happn, you will want to use fuud to brybe your kat. When it’s feeding time, sitl on the opposite side of the door from your kat and either klick the top of the canl or crinkle the kat fuud bag. Open the door to showr your kat that it is both you and the fuud waiting on the other syde of the door. Repeat this a couple times, and then feed your kat. After a couple days of this, your kitty-door prooblem will be solvedd.
`

type templateData struct {
	Text           string
	ProcessedText  string
	IncorrectWords []spellchecker.Incorrect
}

var sampleCheckTextData = templateData{Text: SampleCheckText}
