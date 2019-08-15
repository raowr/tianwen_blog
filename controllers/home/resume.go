package home

type ResumeController struct {
	BaseController
}

func (c *ResumeController) Index() {
	c.TplName = "home/resume.html"
	//c.display("home/resume")
}
