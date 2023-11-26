// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2023 Steadybit GmbH

package extcluster

import (
	"context"
	"github.com/steadybit/discovery-kit/go/discovery_kit_api"
	"github.com/steadybit/discovery-kit/go/discovery_kit_sdk"
	"github.com/steadybit/extension-kit/extbuild"
	"github.com/steadybit/extension-kit/extutil"
	"github.com/steadybit/extension-kubernetes/extconfig"
	"time"
)

type clusterDiscovery struct {
}

var (
	_ discovery_kit_sdk.TargetDescriber = (*clusterDiscovery)(nil)
)

func NewClusterDiscovery() discovery_kit_sdk.TargetDiscovery {
	discovery := &clusterDiscovery{}
	return discovery_kit_sdk.NewCachedTargetDiscovery(
		discovery,
		discovery_kit_sdk.WithRefreshTargetsNow(),
		discovery_kit_sdk.WithRefreshTargetsInterval(context.Background(), 60*time.Second),
	)
}

func (c *clusterDiscovery) Describe() discovery_kit_api.DiscoveryDescription {
	return discovery_kit_api.DiscoveryDescription{
		Id:         ClusterTargetType,
		RestrictTo: extutil.Ptr(discovery_kit_api.LEADER),
		Discover: discovery_kit_api.DescribingEndpointReferenceWithCallInterval{
			CallInterval: extutil.Ptr("5m"),
		},
	}
}

func (c *clusterDiscovery) DescribeTarget() discovery_kit_api.TargetDescription {
	return discovery_kit_api.TargetDescription{
		Id:       ClusterTargetType,
		Label:    discovery_kit_api.PluralLabel{One: "Kubernetes Cluster", Other: "Kubernetes Cluster"},
		Category: extutil.Ptr("Kubernetes"),
		Version:  extbuild.GetSemverVersionStringOrUnknown(),
		Icon:     extutil.Ptr("data:image/svg+xml,%3Csvg%20width%3D%2224%22%20height%3D%2224%22%20viewBox%3D%220%200%2024%2024%22%20fill%3D%22none%22%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%3E%0A%3Cpath%20d%3D%22M11.9481%2013.0696L12.5657%2012.7723L12.7258%2012.109L12.2912%2011.5601H11.6051L11.1705%2012.109L11.3306%2012.7723L11.9481%2013.0696Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20d%3D%22M16.3851%2012.7265C16.4308%2012.2462%2016.408%2011.7659%2016.2936%2011.2857C16.1793%2010.8054%2015.9963%2010.348%2015.7447%209.93627L14.0294%2011.4915C13.9837%2011.5372%2013.9608%2011.583%2013.9379%2011.6516C13.8922%2011.8574%2014.0065%2012.0633%2014.2124%2012.109L16.3851%2012.7265Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20d%3D%22M13.1832%2010.4166L15.0586%209.09005C14.3725%208.40393%2013.4805%207.96938%2012.4742%207.85502L12.6114%2010.165C12.6114%2010.2336%2012.6343%2010.2793%2012.68%2010.3251C12.7944%2010.4852%2013.0231%2010.508%2013.1832%2010.4166Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20d%3D%22M11.3764%207.83215L10.919%207.92364C10.1185%208.1066%209.38661%208.51828%208.79197%209.09005L10.6902%2010.4394C10.7046%2010.4442%2010.7179%2010.449%2010.7306%2010.4535C10.779%2010.4709%2010.8189%2010.4852%2010.8732%2010.4852C11.0791%2010.4852%2011.262%2010.3251%2011.262%2010.1192L11.3764%207.83215Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20d%3D%22M9.82116%2011.4458L8.12871%209.93627C7.62555%2010.7825%207.39684%2011.7659%207.46545%2012.7265L9.68393%2012.0861C9.77541%2012.0633%209.82115%2012.0404%209.8669%2011.9718C10.0041%2011.8117%209.98125%2011.583%209.82116%2011.4458Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20d%3D%22M10.0041%2013.4126L7.7399%2013.8015C8.0601%2014.7163%208.70048%2015.5168%209.50096%2016.0428L10.3701%2013.9387C10.4158%2013.8701%2010.4158%2013.7786%2010.3929%2013.71C10.3701%2013.527%2010.1871%2013.4126%2010.0041%2013.4126Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20d%3D%22M11.9481%2016.7518C12.2912%2016.7518%2012.6114%2016.7061%2012.9316%2016.6374C13.0917%2016.5917%2013.2289%2016.5688%2013.3661%2016.546L12.2683%2014.5562C12.1997%2014.4876%2012.154%2014.4418%2012.0854%2014.3961C11.9253%2014.3046%2011.7423%2014.3504%2011.628%2014.4876L10.5073%2016.5231C10.9647%2016.6603%2011.4679%2016.7518%2011.9481%2016.7518Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20d%3D%22M14.3496%2016.0199C14.8985%2015.6769%2015.3788%2015.1966%2015.7218%2014.6477C15.9048%2014.3961%2016.042%2014.0988%2016.1564%2013.7786L13.8693%2013.3898C13.8007%2013.3898%2013.7321%2013.4126%2013.6635%2013.4355C13.5034%2013.5041%2013.4119%2013.6871%2013.4576%2013.8701L14.3496%2016.0199Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M19.5185%205.4535C19.8616%205.63647%2020.136%205.93379%2020.2504%206.29973L21.9886%2013.7785C22.0343%2014.1673%2021.9428%2014.5561%2021.7141%2014.8763L16.8884%2020.8456C16.6368%2021.1887%2016.248%2021.3716%2015.8363%2021.3259H8.15168C7.76288%2021.303%207.37407%2021.1201%207.09962%2020.8456L2.27386%2014.8763C2.04516%2014.5561%201.95367%2014.1673%202.02228%2013.7785L3.7376%206.25398C3.82908%205.86518%204.08066%205.56786%204.42373%205.40776L11.3993%202.04574C11.5823%202%2011.7882%202%2011.9711%202C12.1541%202%2012.3599%202.02287%2012.5429%202.11435L19.5185%205.4535ZM19.1296%2013.4584C19.1296%2013.4813%2019.1525%2013.4813%2019.1753%2013.4813C19.3812%2013.527%2019.5642%2013.71%2019.587%2013.9615C19.5184%2014.1445%2019.3354%2014.2817%2019.1296%2014.2817C19.0839%2014.2817%2019.061%2014.2817%2019.0153%2014.2589C18.9924%2014.236%2018.9924%2014.236%2018.9695%2014.236C18.9427%2014.236%2018.9238%2014.2282%2018.9081%2014.2217C18.897%2014.2171%2018.8875%2014.2131%2018.878%2014.2131C18.8102%2014.1962%2018.7549%2014.1666%2018.6936%2014.1338C18.6722%2014.1224%2018.6501%2014.1106%2018.6264%2014.0988C18.5807%2014.0988%2018.535%2014.0759%2018.4892%2014.053H18.4664C18.2376%2013.9615%2017.9861%2013.8929%2017.7345%2013.8472H17.7116C17.643%2013.8472%2017.5744%2013.8701%2017.5286%2013.9158C17.5286%2013.9158%2017.5286%2013.9387%2017.5058%2013.9387L17.3228%2013.9158C16.9111%2015.2194%2016.0192%2016.3173%2014.8528%2017.0262L14.9214%2017.2092C14.9214%2017.2092%2014.8985%2017.2092%2014.8985%2017.2321C14.8527%2017.3007%2014.8527%2017.3922%2014.8756%2017.4608C14.9671%2017.6895%2015.0815%2017.9182%2015.2416%2018.1241V18.1698C15.2873%2018.2155%2015.3102%2018.2384%2015.333%2018.2841C15.4017%2018.3528%2015.4474%2018.4214%2015.4931%2018.5129C15.516%2018.5357%2015.5389%2018.5586%2015.5389%2018.5815C15.5389%2018.5815%2015.5617%2018.5815%2015.5617%2018.6043C15.6075%2018.7187%2015.6075%2018.833%2015.5846%2018.9474C15.5617%2019.0618%2015.4703%2019.1532%2015.3788%2019.199C15.3645%2019.2038%2015.3511%2019.2085%2015.3384%2019.2131C15.2901%2019.2304%2015.2501%2019.2447%2015.1958%2019.2447C15.0128%2019.2447%2014.8527%2019.1304%2014.7613%2018.9703C14.7384%2018.9703%2014.7384%2018.9474%2014.7384%2018.9474C14.727%2018.936%2014.7212%2018.9245%2014.7155%2018.9131C14.7098%2018.9017%2014.7041%2018.8902%2014.6927%2018.8788C14.6469%2018.8102%2014.624%2018.7187%2014.6012%2018.6272L14.5554%2018.49V18.4671C14.4868%2018.2155%2014.3725%2017.9868%2014.2581%2017.7581C14.2124%2017.6895%2014.1438%2017.6438%2014.0751%2017.6209C14.0751%2017.598%2014.0751%2017.598%2014.0523%2017.598L13.9608%2017.4379C13.9026%2017.4554%2013.8429%2017.4743%2013.7821%2017.4936C13.604%2017.5501%2013.4165%2017.6097%2013.2289%2017.6438C12.8172%2017.7581%2012.4056%2017.8039%2011.9939%2017.8039C11.3078%2017.8039%2010.6445%2017.6895%2010.0041%2017.4379L9.91264%2017.6209C9.91264%2017.6323%209.91264%2017.638%209.90978%2017.6409C9.90692%2017.6438%209.9012%2017.6438%209.88977%2017.6438C9.82116%2017.6666%209.75254%2017.7124%209.7068%2017.781C9.59245%2018.0097%209.47809%2018.2384%209.40948%2018.49L9.36374%2018.6272C9.3523%2018.673%209.33515%2018.713%209.318%2018.753C9.30084%2018.793%209.28369%2018.833%209.27225%2018.8788C9.24938%2018.9017%209.22651%2018.9245%209.22651%2018.9474C9.20364%2018.9474%209.20364%2018.9703%209.20364%2018.9703C9.11216%2019.1304%208.95206%2019.2447%208.76909%2019.2447C8.72335%2019.2447%208.65474%2019.2219%208.609%2019.199C8.40316%2019.0846%208.31168%2018.8102%208.42603%2018.5815C8.4489%2018.5815%208.4489%2018.5586%208.4489%2018.5586C8.46034%2018.5472%208.46605%2018.5357%208.47177%2018.5243C8.47749%2018.5129%208.48321%2018.5014%208.49464%2018.49C8.52473%2018.4499%208.55041%2018.4098%208.57363%2018.3735C8.60337%2018.3271%208.62905%2018.287%208.65474%2018.2613C8.70048%2018.2155%208.72335%2018.1927%208.74622%2018.1469V18.1241C8.88345%2017.9182%209.02067%2017.6895%209.11216%2017.4608C9.13503%2017.3922%209.13503%2017.3007%209.08929%2017.2321C9.08929%2017.2321%209.06642%2017.2321%209.06642%2017.2092L9.18077%2017.0491C8.95206%2016.9348%208.76909%2016.7975%208.56326%2016.6374C7.67129%2015.9513%207.03091%2015.0136%206.68784%2013.9615L6.48201%2013.9844C6.48201%2013.9844%206.48201%2013.9615%206.45914%2013.9615C6.41339%2013.9158%206.34478%2013.8929%206.27617%2013.8929H6.2533C5.97885%2013.9387%205.75014%2014.0073%205.49856%2014.0988H5.47569C5.42994%2014.0988%205.3842%2014.1216%205.33846%2014.1445C5.31292%2014.153%205.2842%2014.1647%205.2535%2014.1772C5.20173%2014.1983%205.14431%2014.2216%205.08688%2014.236C5.07996%2014.236%205.06884%2014.2339%205.0567%2014.2316C5.02875%2014.2263%204.9954%2014.2201%204.9954%2014.236C4.9954%2014.2474%204.9954%2014.2532%204.99254%2014.256C4.98968%2014.2589%204.98396%2014.2589%204.97253%2014.2589C4.92679%2014.2817%204.90391%2014.2817%204.85817%2014.2817C4.65233%2014.3046%204.4465%2014.1674%204.40075%2013.9615C4.35501%2013.71%204.51511%2013.4813%204.76669%2013.4355C4.77984%2013.4224%204.78543%2013.4168%204.79215%2013.4144C4.79712%2013.4126%204.80271%2013.4126%204.81243%2013.4126C4.83923%2013.4126%204.85817%2013.4048%204.87387%2013.3983C4.88497%2013.3937%204.89444%2013.3898%204.90391%2013.3898C4.94966%2013.3898%204.9954%2013.3841%205.04114%2013.3783C5.08688%2013.3726%205.13262%2013.3669%205.17837%2013.3669C5.22411%2013.344%205.26985%2013.344%205.31559%2013.344C5.59004%2013.3212%205.84162%2013.2754%206.0932%2013.2068C6.16181%2013.1611%206.23043%2013.1153%206.2533%2013.0467C6.2533%2013.0467%206.27617%2013.0467%206.27617%2013.0238L6.45914%2012.9781C6.2533%2011.6745%206.55062%2010.348%207.25962%209.22728C7.27105%209.20441%207.28249%209.18725%207.29392%209.1701C7.30536%209.15295%207.31679%209.13579%207.32823%209.11292L7.19546%208.98016C7.2059%208.91436%207.16415%208.83449%207.12239%208.79273C6.93942%208.60976%206.71071%208.47254%206.48201%208.33531L6.34478%208.2667C6.2533%208.22096%206.16181%208.17522%206.0932%208.12947C6.07033%208.12947%206.02459%208.08373%206.02459%208.08373C6.02459%208.08373%206.02459%208.06086%206.00172%208.06086C5.81875%207.90077%205.77301%207.62632%205.91023%207.42048C5.97885%207.30612%206.0932%207.26038%206.23043%207.26038C6.34478%207.26038%206.45914%207.30612%206.55062%207.37474L6.57349%207.39761C6.58492%207.40904%206.59636%207.41476%206.6078%207.42048C6.61923%207.4262%206.63067%207.43191%206.6421%207.44335C6.67641%207.47765%206.705%207.51196%206.73358%207.54627C6.76217%207.58057%206.79076%207.61488%206.82507%207.64919C6.83199%207.65611%206.84102%207.66304%206.85086%207.6706C6.87354%207.688%206.90061%207.70878%206.91655%207.74067C7.07665%207.92364%207.28249%208.1066%207.48832%208.2667C7.53407%208.28957%207.57981%208.31244%207.62555%208.31244C7.65235%208.31244%207.67129%208.30459%207.68699%208.29809C7.69809%208.29349%207.70756%208.28957%207.71703%208.28957H7.7399L7.87713%208.38105C8.63187%207.58057%209.61532%207.0088%2010.6902%206.78009C10.7299%206.77349%2010.769%206.76689%2010.8077%206.76035C11.0373%206.72162%2011.2526%206.68531%2011.4679%206.66574L11.4907%206.48277V6.43703C11.5593%206.39129%2011.5822%206.32268%2011.6051%206.25406C11.6051%205.97961%2011.6051%205.72803%2011.5593%205.47645V5.45358C11.5593%205.40784%2011.5593%205.3621%2011.5365%205.31636C11.5136%205.22487%2011.4907%205.13339%2011.4907%205.04191V4.92755C11.4907%204.8132%2011.5365%204.69884%2011.628%204.60736C11.7423%204.493%2011.8795%204.42439%2012.0168%204.44726C12.2455%204.47013%2012.4284%204.69884%2012.4056%204.92755V5.06478C12.3941%205.11052%2012.3884%205.15626%2012.3827%205.202C12.377%205.24774%2012.3713%205.29349%2012.3598%205.33923C12.3598%205.3621%2012.3541%205.38497%2012.3484%205.40784C12.3427%205.43071%2012.337%205.45358%2012.337%205.47645V5.49932C12.2912%205.77377%2012.2912%206.02535%2012.2912%206.27693C12.3141%206.34555%2012.337%206.41416%2012.4056%206.4599V6.43703L12.4284%206.62C13.5034%206.73435%2014.5554%207.14603%2015.4017%207.80928C15.4931%207.90077%2015.5903%207.99225%2015.6875%208.08373C15.7847%208.17522%2015.8819%208.2667%2015.9734%208.35818L16.1564%208.24383H16.1793C16.2021%208.2667%2016.2479%208.2667%2016.2707%208.2667C16.3165%208.2667%2016.3622%208.24383%2016.408%208.22096C16.6138%208.08373%2016.8196%207.90077%2016.9797%207.7178C16.9867%207.71087%2016.9957%207.70395%2017.0055%207.69639C17.0282%207.67898%2017.0553%207.65821%2017.0712%207.62632C17.117%207.55771%2017.1856%207.4891%2017.2542%207.42049C17.2771%207.42049%2017.2999%207.39761%2017.3228%207.37474L17.3457%207.35186C17.4372%207.28325%2017.5515%207.23751%2017.6659%207.23751C17.7802%207.23751%2017.9174%207.30612%2017.9861%207.39761C18.1462%207.60344%2018.1004%207.87789%2017.8946%208.03799C17.8946%208.04673%2017.8979%208.05213%2017.9008%208.05674C17.9054%208.0642%2017.9087%208.0696%2017.8946%208.08373C17.8831%208.09517%2017.8717%208.10089%2017.8603%208.1066C17.8488%208.11232%2017.8374%208.11804%2017.826%208.12947C17.7802%208.15235%2017.7402%208.17522%2017.7002%208.19809C17.6601%208.22096%2017.6201%208.24383%2017.5744%208.2667L17.4372%208.33531C17.2085%208.47254%2017.0026%208.60976%2016.7968%208.79273C16.751%208.83847%2016.7282%208.92996%2016.7282%208.99857V9.02144L16.5909%209.15866C16.9569%209.73044%2017.2313%2010.3708%2017.3914%2011.0341C17.5286%2011.6973%2017.5744%2012.3835%2017.4829%2013.0467L17.6659%2013.0925C17.6887%2013.1611%2017.7574%2013.2297%2017.826%2013.2526C18.0775%2013.3212%2018.352%2013.3669%2018.6036%2013.3898H18.6264C18.6722%2013.4126%2018.7179%2013.4126%2018.7637%2013.4126C18.8552%2013.4126%2018.9466%2013.4126%2019.0381%2013.4355C19.0839%2013.4355%2019.1296%2013.4355%2019.1296%2013.4584ZM7.19546%208.98016L7.191%208.9757V8.99857C7.193%208.99259%207.19447%208.98644%207.19546%208.98016Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3C%2Fsvg%3E"),
		Table: discovery_kit_api.Table{
			Columns: []discovery_kit_api.Column{
				{Attribute: "k8s.cluster-name"},
			},
			OrderBy: []discovery_kit_api.OrderBy{
				{
					Attribute: "k8s.cluster-name",
					Direction: "ASC",
				},
			},
		},
	}
}

func (c *clusterDiscovery) DiscoverTargets(_ context.Context) ([]discovery_kit_api.Target, error) {
	return []discovery_kit_api.Target{
		{
			Id:         extconfig.Config.ClusterName,
			Label:      extconfig.Config.ClusterName,
			TargetType: ClusterTargetType,
			Attributes: map[string][]string{
				"k8s.cluster-name": {extconfig.Config.ClusterName},
			},
		},
	}, nil
}
