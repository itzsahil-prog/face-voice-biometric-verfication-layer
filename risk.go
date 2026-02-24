package engine

import (
	"math"
	"fintech-app/internal/models"
)

// RiskFactor weights
const (
	WeightAmount      = 0.4
	WeightLocation    = 0.2
	WeightDevice      = 0.2
	WeightBiometric   = 0.2
)

type RiskEngine struct{}

func NewRiskEngine() *RiskEngine {
	return &RiskEngine{}
}

type RiskDecision struct {
	Score       int      `json:"risk_score"` // 0-100
	Level       string   `json:"level"`      // LOW, MEDIUM, HIGH
	RequiredAuth []string `json:"required_auth"` // FACE, VOICE, OTP
}

func (re *RiskEngine) CalculateRisk(txn *models.Transaction, user *models.User, bioConfidence float64) RiskDecision {
	score := 0
	factors := []string{}

	// 1. Amount Risk
	if txn.Amount > 10000 {
		score += int(WeightAmount * 100)
		factors = append(factors, "high_amount")
	} else if txn.Amount > 1000 {
		score += int(WeightAmount * 50)
	}

	// 2. Biometric Confidence (The Core Trust Layer)
	// If biometric failed or confidence is low (< 0.85), spike risk
	if bioConfidence < 0.85 {
		score += 50
		factors = append(factors, "low_bio_confidence")
	}

	// 3. Device Fingerprint (Simplified)
	if user.DeviceFingerprint == "" {
		score += 20
		factors = append(factors, "unknown_device")
	}

	// Determine Required Auth based on final score
	requiredAuth := []string{"FACE"} // Default
	
	if score > 75 {
		return RiskDecision{
			Score:       score,
			Level:       "HIGH",
			RequiredAuth: []string{"FACE", "