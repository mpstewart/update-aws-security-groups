package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Singleton for EC2 instance
var svc *ec2.EC2

func GetSvc() (s *ec2.EC2) {
	config := GetConfig()
	region := config.Region
	if svc == nil {
		svc = ec2.New(
			session.Must(
				session.NewSession(
					&aws.Config{Region: aws.String(region)},
				),
			),
		)
	}

	s = svc

	return
}

// Given a port, call Amazon and ask them to update the security rules for that
// port
func UpdateRuleForPort(p Port) {
	input := AuthorizeSecurityGroupIngressInput(p)

	amazonSvc := GetSvc()

	_, err := amazonSvc.AuthorizeSecurityGroupIngress(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case "InvalidPermission.Duplicate":
				// If this triggers, nothing is actually wrong
				Logger.Printf("Duplicate permission found for Port %v\n", p)
			default:
				Logger.Println(aerr.Error())
			}
		} else {
			// It wasn't a special awserr.Error, but we can still just log the output
			// of the Error() function
			Logger.Println(err.Error())
		}
	}
}

// Amazon's AWS SDK uses a special struct to track the input arguments for each
// of its functions, which is a little weird, but this method generates that for
// a given port.
func AuthorizeSecurityGroupIngressInput(p Port) *ec2.AuthorizeSecurityGroupIngressInput {
	config := GetConfig()
	ipa := config.HomeIP
	groupID := config.GroupID

	return &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: aws.String(groupID),
		IpPermissions: []*ec2.IpPermission{
			{
				FromPort:   aws.Int64(p.Port),
				IpProtocol: aws.String(p.Protocol),
				IpRanges: []*ec2.IpRange{
					{
						CidrIp:      aws.String(ipa),
						Description: aws.String(p.Description),
					},
				},
				ToPort: aws.Int64(p.Port),
			},
		},
	}
}
