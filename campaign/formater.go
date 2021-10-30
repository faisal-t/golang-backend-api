package campaign

import "strings"

type CampaignFormater struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

//for formating response
func FormatCampaign(campaign Campaign) CampaignFormater {
	campaignFormater := CampaignFormater{}
	campaignFormater.ID = campaign.ID
	campaignFormater.UserID = campaign.UserID
	campaignFormater.Name = campaign.Name
	campaignFormater.ShortDescription = campaign.ShortDescription
	campaignFormater.GoalAmount = campaign.GoalAmount
	campaignFormater.CurrentAmount = campaign.CurrentAmount
	campaignFormater.Slug = campaign.Slug
	campaignFormater.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormater.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return campaignFormater
}

// format campaigns
func FormatCampaigns(campaigns []Campaign) []CampaignFormater {

	campaignsFormater := []CampaignFormater{}

	for _, campaign := range campaigns {
		campaignFormater := FormatCampaign(campaign)
		campaignsFormater = append(campaignsFormater, campaignFormater)
	}

	return campaignsFormater

}

type CampaignDetailFormater struct {
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	ShortDescription string   `json:"short_description"`
	Description      string   `json:"description"`
	ImageUrl         string   `json:"image_url"`
	GoalAmount       int      `json:"goal_amount"`
	CurrentAmount    int      `json:"current_amount"`
	UserID           int      `json:"user_id"`
	Slug             string   `json:"slug"`
	Perks            []string `json:"perks"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormater {
	campaiDetailgnFormater := CampaignDetailFormater{}

	campaiDetailgnFormater.ID = campaign.ID
	campaiDetailgnFormater.Name = campaign.Name
	campaiDetailgnFormater.Description = campaign.Description
	campaiDetailgnFormater.ShortDescription = campaign.ShortDescription
	campaiDetailgnFormater.GoalAmount = campaign.GoalAmount
	campaiDetailgnFormater.CurrentAmount = campaign.CurrentAmount
	campaiDetailgnFormater.Slug = campaign.Slug
	campaiDetailgnFormater.ImageUrl = ""
	campaiDetailgnFormater.UserID = campaign.UserID

	if len(campaign.CampaignImages) > 0 {
		campaiDetailgnFormater.ImageUrl = campaign.CampaignImages[0].FileName
	}

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaiDetailgnFormater.Perks = perks

	return campaiDetailgnFormater
}
